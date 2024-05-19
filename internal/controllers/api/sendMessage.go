package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Payload struct {
	ChatID           string `json:"chatId"`
	Message          string `json:"message"`
	IdInstance       string `json:"idInstance" binding:"required"`
	ApiTokenInstance string `json:"apiTokenInstance" binding:"required"`
}

func SendMessage(c *gin.Context) {
	var payload Payload
	if err := c.ShouldBindJSON(&payload); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiUrl := "https://api.greenapi.com"
	if len(payload.IdInstance) >= 4 {
		apiUrl = fmt.Sprintf("https://%s.api.greenapi.com", payload.IdInstance[:4])
	}

	url := fmt.Sprintf("%s/waInstance%s/sendMessage/%s", apiUrl, payload.IdInstance, payload.ApiTokenInstance)

	data := struct {
		ChatID  string `json:"chatId"`
		Message string `json:"message"`
	}{}
	data.ChatID = payload.ChatID
	data.Message = payload.Message

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), respBody)
}
