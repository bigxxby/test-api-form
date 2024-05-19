package controllers

import "github.com/gin-gonic/gin"

func GET_IndexHTML(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
