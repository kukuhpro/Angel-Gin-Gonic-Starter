package main

import (
	"gopkg.in/gin-gonic/gin"
	"fmt"
	"./config"
)


func resPing(c *gin.Context) {
	c.JSON(200, gin.H {
		"message": "Hello World!",
	})
}

func main() {
	list := map[string]interface{} {
		"GET": "ping",
	}
    r := gin.Default()
	for key,val := range list {
		if fmt.Sprint(key) == "GET" {
			r.GET(fmt.Sprint(val), resPing)
		}
	}
    r.Run(":3214") 
}