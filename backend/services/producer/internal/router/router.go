package router

import (
	"net/http"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/domain"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/producer/internal/gameevent"
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

	// Get user value
	r.POST("/produce", func(c *gin.Context) {
		
		status, body, err := gameevent.HandleGameEvent(c.Request.Body)

		if err != nil{
			 c.JSON(status, &domain.ApiResponse{
				Status: uint32(status),
				Message: err.Error(),
			 })
		}

		c.JSON(status, &domain.ApiResponse{
				Status: 200,
				Data: body,
			 })

	})

	return r
}
