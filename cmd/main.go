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
	router.Run(":8080")

}
func HelloWorld(c *gin.Context) {
	c.JSON(200, "Hello world")
}
func GetEnv(c *gin.Context) {
	env := os.Getenv("TEST_ENV")

	c.JSON(200, "test env : "+env)

}
