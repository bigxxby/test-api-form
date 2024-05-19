package main

import (
	"log"
	"os"
	"project/internal/controllers"
	"project/internal/controllers/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./ui/static")
	router.LoadHTMLGlob("ui/templates/*")

	//ROUTES

	//html
	router.GET("/index", controllers.GET_IndexHTML)
	router.POST("/api/getSettings", api.GetSettings)
	router.POST("/api/getStateInstance", api.GetStateInstance)
	router.POST("/api/sendMessage", api.SendMessage)
	router.POST("/api/sendFileByUrl", api.SendFileByUrl)

	//ROUTES

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started")
	err := router.Run(":" + port)
	if err != nil {
		log.Println(err.Error())
		return
	}

}
