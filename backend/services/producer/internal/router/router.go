package router

import (
	"net/http"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/domain"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/producer/internal/gameevent"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/producer/internal/kafka"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Read store
	r.GET("/read", func(c *gin.Context) {
		playerEvents,sessionEvents, shipEvents := kafka.Read()
		data := make(map[string]interface{})
		data["player"] = playerEvents
		data["session"] = sessionEvents
		data["ship"] = shipEvents

		c.JSON(http.StatusOK, &domain.ApiResponse{
				Status: uint32(http.StatusOK),
				Message: domain.OK,
				Data: data,
			})
	})

	// Get user value
	r.POST("/produce", func(c *gin.Context) {

		status, body, err2 := gameevent.HandleGameEvent(c)

		if err2 != nil{
			c.JSON(status, &domain.ApiResponse{
				Status: uint32(status),
				Message: err2.Error(),
			})

			return
		}

		c.JSON(status, &domain.ApiResponse{
				Status: 200,
				Data: body,
			 })

	})

	return r
}
