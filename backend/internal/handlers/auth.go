package handlers

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "github.com/C14147/SmartCampus-Workbench/internal/config"
    "github.com/C14147/SmartCampus-Workbench/internal/models"
    "github.com/C14147/SmartCampus-Workbench/internal/utils"
    "github.com/C14147/SmartCampus-Workbench/pkg/response"
)

type loginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type registerRequest struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

// RegisterHandler creates a user (uses GORM via context)
func RegisterHandler(c *gin.Context) {
    var req registerRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "invalid request", err.Error())
        return
    }

    if err := utils.ValidateStruct(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "validation failed", err.Error())
        return
    }

    db, ok := c.Get("db")
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database not configured"})
        return
    }
    gdb := db.(*gorm.DB)

    hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    user := &models.User{Username: req.Username, Email: req.Email, PasswordHash: string(hash), Role: "student"}
    if err := gdb.Create(user).Error; err != nil {
        response.Error(c, http.StatusBadRequest, "create user failed", err.Error())
        return
    }

    response.Success(c, gin.H{"id": user.ID, "username": user.Username})
}

// LoginHandler verifies credentials and returns JWT
func LoginHandler(c *gin.Context) {
    var req loginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "invalid request", err.Error())
        return
    }

    if err := utils.ValidateStruct(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "validation failed", err.Error())
        return
    }

    db, ok := c.Get("db")
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database not configured"})
        return
    }
    gdb := db.(*gorm.DB)

    var user models.User
    if err := gdb.Where("username = ?", req.Username).First(&user).Error; err != nil {
        response.Error(c, http.StatusUnauthorized, "invalid credentials", nil)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        response.Error(c, http.StatusUnauthorized, "invalid credentials", nil)
        return
    }

    cfg, _ := config.LoadConfig()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub":  user.ID,
        "name": user.Username,
        "role": user.Role,
        "exp":  time.Now().Add(time.Hour * 24).Unix(),
    })
    signed, err := token.SignedString([]byte(cfg.JWT.Secret))
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "token generation failed", err.Error())
        return
    }

    response.Success(c, gin.H{"token": signed})
}

// MeHandler returns a minimal current user; the AuthMiddleware will set userID in context
func MeHandler(c *gin.Context) {
    uid, ok := c.Get("user_id")
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
        return
    }
    db, _ := c.Get("db")
    gdb := db.(*gorm.DB)

    var user models.User
    if err := gdb.First(&user, "id = ?", uid).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    response.Success(c, gin.H{"id": user.ID, "username": user.Username, "role": user.Role})
}

// AuthMiddleware parses JWT and sets user_id in context
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if auth == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
            return
        }
        // expect Bearer <token>
        var tokenString string
        fmt.Sscanf(auth, "Bearer %s", &tokenString)

        cfg, _ := config.LoadConfig()
        token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
            return []byte(cfg.JWT.Secret), nil
        })
        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }
        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            if sub, ok := claims["sub"].(string); ok {
                c.Set("user_id", sub)
            }
            if role, ok := claims["role"].(string); ok {
                c.Set("user_role", role)
            }
        }
        c.Next()
    }
}

