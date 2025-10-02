package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
    c.JSON(200, gin.H{"success": true, "data": data})
}

func BadRequest(c *gin.Context, msg string) {
    c.JSON(400, gin.H{"success": false, "error": msg})
}

func Unauthorized(c *gin.Context, msg string) {
    c.JSON(401, gin.H{"success": false, "error": msg})
}

func Forbidden(c *gin.Context, msg string) {
    c.JSON(403, gin.H{"success": false, "error": msg})
}

func ServerError(c *gin.Context, msg string) {
    c.JSON(500, gin.H{"success": false, "error": msg})
}
