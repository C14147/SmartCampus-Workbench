package auth

import (
    "net/http"

    "github.com/casbin/casbin/v2"
    "github.com/gin-gonic/gin"
)

// RequirePermission checks Casbin policy for current user role and requested path
func RequirePermission(e *casbin.Enforcer) gin.HandlerFunc {
    return func(c *gin.Context) {
        roleIfc, _ := c.Get("user_role")
        role, _ := roleIfc.(string)
        if role == "" {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "no role"})
            return
        }

        obj := c.FullPath()
        act := c.Request.Method

        ok, err := e.Enforce(role, obj, act)
        if err != nil || !ok {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
            return
        }
        c.Next()
    }
}
