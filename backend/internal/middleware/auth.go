package middleware

import (
    "github.com/gin-gonic/gin"
    "smartcampus/internal/utils"
    "smartcampus/pkg/response"
)

func AuthMiddleware(jwtUtil utils.JWTUtil) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            response.Unauthorized(c, "缺少认证token")
            c.Abort()
            return
        }

        claims, err := jwtUtil.ValidateToken(tokenString)
        if err != nil {
            response.Unauthorized(c, "无效的token")
            c.Abort()
            return
        }

        c.Set("userID", claims.UserID)
        c.Set("userRole", claims.Role)
        c.Next()
    }
}

func RBAC(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("userRole")
        if !exists {
            response.Unauthorized(c, "未认证用户")
            c.Abort()
            return
        }

        for _, role := range allowedRoles {
            if userRole == role {
                c.Next()
                return
            }
        }

        response.Forbidden(c, "权限不足")
        c.Abort()
    }
}
