package main

import (
	"github.com/GreedHub/battleship-event-sourcing/backend/services/producer/internal/router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}
