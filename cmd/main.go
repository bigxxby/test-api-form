package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", HelloWorld)
	router.GET("/env", GetEnv)
	log.Println("Server started")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	err := router.Run(":" + port)
	if err != nil {
		log.Println(err.Error())
		return
	}

}
func HelloWorld(c *gin.Context) {
	c.JSON(200, "Hello world")
}
func GetEnv(c *gin.Context) {
	env := os.Getenv("TEST_ENV")

	c.JSON(200, "test env : "+env)

}
