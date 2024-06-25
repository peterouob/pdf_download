package main

import (
	"file_download/config"
	controllers "file_download/controller"
	"file_download/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	config.LoadConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr: config.Cfg.RedisAddr,
	})
	defer rdb.Close()
	downloadService := service.NewDownloadService(rdb)
	downloadController := controllers.NewDownloadController(downloadService)

	r := gin.Default()
	r.POST("/", downloadController.Download)
	r.Run(config.Cfg.Addr)
}
