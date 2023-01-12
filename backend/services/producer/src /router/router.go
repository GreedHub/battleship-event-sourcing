package router

import (
	"net/http"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/src/domain"
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
	r.POST("/EventHandler", func(c *gin.Context) {
		
		err := gameevent.HandleGameEvent(c.Request.Body)

		if err != nil{
			 c.JSON(500, &domain.ApiResponse{
				Status: 500,
				Message: domain.INTERNAL_SERVER_ERROR,
			 })
		}

		c.JSON(200, &domain.ApiResponse{
				Status: 200,
				Message: domain.OK,
			 })


	})

	return r
}
