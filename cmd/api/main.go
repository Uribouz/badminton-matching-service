package main

import (
	"badminton-service/common/cache"
	"badminton-service/common/database"
	"badminton-service/common/logger"
	"badminton-service/common/middleware"
	"badminton-service/config"
	"badminton-service/domain/device"
	"badminton-service/domain/event"
	"badminton-service/domain/match"
	"badminton-service/domain/partner"
	"badminton-service/domain/player"
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
		AllowOrigins: config.App.AllowOrigins, //TODO: make this secure
	}))
	var (
		deviceRepo    = device.NewMongoRepository(database.Database)
		deviceService = device.NewService(deviceRepo)
		deviceHandler = device.NewHandler(deviceService)

		eventRepo    = event.NewMongoRepository(database.Database)
		eventService = event.NewService(eventRepo)
		eventHandler = event.NewHandler(eventService)

		playerRepo    = player.NewMongoRepository(database.Database)
		playerService = player.NewService(playerRepo)
		playerHandler = player.NewHandler(playerService)

		matchRepo    = match.NewMongoRepository(database.Database)
		matchService = match.NewService(matchRepo)
		matchHandler = match.NewHandler(matchService)

		partnerRepo    = partner.NewMongoRepository(database.Database)
		partnerService = partner.NewService(partnerRepo)
		partnerHandler = partner.NewHandler(partnerService)
	)

	deviceGroup := r.Group("/devices")
	deviceGroup.GET("/:id", deviceHandler.Get)
	deviceGroup.POST("", deviceHandler.Post)

	eventGroup := r.Group("/events")
	eventGroup.GET("/:event_id", eventHandler.Get)
	eventGroup.POST("", eventHandler.Post)

	playerGroup := r.Group("/players")
	playerGroup.GET("/:event_id/:player_name", playerHandler.Get)
	playerGroup.POST("", playerHandler.Post)

	matchGroup := r.Group("/matches")
	matchGroup.GET("/:event_id/:court_no/:date_time", matchHandler.Get)
	matchGroup.POST("", matchHandler.Post)

	partnerGroup := r.Group("/partners")
	partnerGroup.GET("/:event_id/:player_name", partnerHandler.Get)
	partnerGroup.POST("", partnerHandler.Post)

	r.Run(":8080")
}
