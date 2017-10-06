package main

import "github.com/gin-gonic/gin"
import "github.com/tustak/moozo/auth"

var customPort string = ":3030"

func main() {
    r := gin.Default()

    // statics
    r.Static("/assets", "./assets")

    // routing
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.GET("login", auth.LoginHandle)
    r.Run(customPort) // listen and serve on 0.0.0.0:8080
}


