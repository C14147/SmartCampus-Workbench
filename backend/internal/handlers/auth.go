package handlers

import (
    "github.com/gin-gonic/gin"
    "smartcampus/internal/services"
    "smartcampus/internal/utils"
    "smartcampus/pkg/response"
)

type AuthHandler struct {
    UserService services.UserService
    JWTUtil     utils.JWTUtil
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Token string      `json:"token"`
    User  interface{} `json:"user"`
}

func (h *AuthHandler) Login(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.BadRequest(c, "无效的请求参数")
        return
    }

    user, err := h.UserService.Authenticate(req.Username, req.Password)
    if err != nil {
        response.Unauthorized(c, "用户名或密码错误")
        return
    }

    token, err := h.JWTUtil.GenerateToken(user.ID, user.Role)
    if err != nil {
        response.ServerError(c, "生成token失败")
        return
    }

    response.Success(c, LoginResponse{
        Token: token,
        User:  user.ToDTO(),
    })
}
