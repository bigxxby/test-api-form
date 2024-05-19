package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStateInstance(c *gin.Context) {
	var requestData RequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiUrl := "https://api.greenapi.com"
	if len(requestData.IdInstance) >= 4 {
		apiUrl = fmt.Sprintf("https://%s.api.greenapi.com", requestData.IdInstance[:4])
	}

	url := fmt.Sprintf("%s/waInstance%s/getStateInstance/%s", apiUrl, requestData.IdInstance, requestData.ApiTokenInstance)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
