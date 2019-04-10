package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xTobu/tatara-co/services/api/models"
	"github.com/xTobu/tatara-co/services/api/packages/helpers"
	"github.com/xTobu/tatara-co/services/api/packages/setting"
)

// ENV 取得環境變數 -ldflags "-X main.ENV=production"
// 改從 os 下給 Environment variables, 比較直覺
var ENV string

func main() {

	serverMode := os.Getenv("ENV")
	fmt.Println(serverMode)

	setting.Setup()
	models.Setup()

	router := gin.Default()

	api := router.Group("/api")
	api.Use()
	{
		api.GET("/test", func(c *gin.Context) {
			users, err := models.GetUsers()
			if err != nil {
				log.Fatalf("models.Setup err: %v", err)
			}
			helpers.OutputGormResult(users)
			c.JSON(200, gin.H{
				"status": "success",
				"data":   users,
			})
		})
		api.GET("/test/2", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"status": 200,
			})
		})
	}

	router.Run(":3000")
}
