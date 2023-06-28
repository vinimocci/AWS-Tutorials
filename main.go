package main

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	_ 	 "github.com/go-sql-driver/mysql"
)

func main() {
	routes := gin.Default()
	routes.SetTrustedProxies([]string{})

	routes.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowMethods:     []string{"GET", "POST"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.GET("/hello", func(c *gin.Context) {		
		c.JSON(http.StatusOK, gin.H{
			"message":  "World! =D",
		})
	})

	routes.Run(":5001")
}