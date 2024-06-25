package controllers

import (
	"file_download/service"
	"github.com/google/uuid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DownloadController struct {
	service *service.DownloadService
}

func NewDownloadController(service *service.DownloadService) *DownloadController {
	return &DownloadController{service: service}
}

func (ctrl *DownloadController) Download(c *gin.Context) {
	url := c.PostForm("url")
	name := uuid.NewString()
	filePath := "./" + name + ".pdf"

	err := ctrl.service.DownloadAndSavePDF(url, filePath, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File downloaded and saved", "url": url})
}
