package main

import (
	. "./DBConnections"
	. "./config"
	"github.com/gin-gonic/gin"
)

func main() {
	LoadConfig()

	router := gin.Default()
	userGroup := router.Group("/CRUD")
	{
		userGroup.POST("/insert", DbInsert)
		//userGroup.POST("/delete", DbDelete)
		userGroup.POST("/read", DbRead)
		//userGroup.POST("/update", DbUpdate)
	}
	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"message": "welcome to CRUD",
		})
	})
	router.Run(Config.PORT)
}
