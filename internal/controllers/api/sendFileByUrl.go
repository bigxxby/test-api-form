package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileRequestData struct {
	IdInstance       string `json:"idInstance" binding:"required"`
	ApiTokenInstance string `json:"apiTokenInstance" binding:"required"`
	ChatId           string `json:"chatId" binding:"required"`
	UrlFile          string `json:"urlFile" binding:"required"`
	FileName         string `json:"fileName" binding:"required"`
}

func SendFileByUrl(c *gin.Context) {
	var requestData FileRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	apiUrl := "https://api.greenapi.com"
	if len(requestData.IdInstance) >= 4 {
		apiUrl = fmt.Sprintf("https://%s.api.greenapi.com", requestData.IdInstance[:4])
	}

	url := fmt.Sprintf("%s/waInstance%s/sendFileByUrl/%s", apiUrl, requestData.IdInstance, requestData.ApiTokenInstance)
	data := struct {
		ChatId   string `json:"chatId" binding:"required"`
		UrlFile  string `json:"urlFile" binding:"required"`
		FileName string `json:"fileName" binding:"required"`
	}{}
	log.Println(url)
	data.ChatId = requestData.ChatId
	data.UrlFile = requestData.UrlFile
	data.FileName = requestData.FileName
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)

}
