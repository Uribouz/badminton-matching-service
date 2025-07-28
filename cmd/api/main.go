package main

import (
	"badminton-service/common/cache"
	"badminton-service/common/database"
	"badminton-service/common/logger"
	"badminton-service/common/middleware"
	"badminton-service/config"
	"badminton-service/domain/device"
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		config          = config.InitConfig()
		logger          = logger.InitLogger()
		_               = cache.NewClient(config.Redis.Host)
		database, dbErr = database.NewMongoDB(config.Mongo)
	)
	if dbErr != nil {
		panic(dbErr)
	}
	defer database.Disconnect(context.Background())

	r := gin.Default()
	r.Use(middleware.LoggingMiddleWare(logger))
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://foo.com"},
	}))
	var (
		deviceRepo    = device.NewMongoRepository(database.Database)
		deviceService = device.NewService(deviceRepo)
		deviceHandler = device.NewHandler(deviceService)
	)
	deviceGroup := r.Group("/device")
	deviceGroup.GET("/:id", deviceHandler.Get)
	deviceGroup.POST("", deviceHandler.Post)

	r.Run(":8080")
}
