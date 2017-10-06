package auth

import "github.com/gin-gonic/gin"

func LoginHandle(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "You are logged in",
    })
}
