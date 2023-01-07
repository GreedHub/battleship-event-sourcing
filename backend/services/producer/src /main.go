package main

import (
	"fmt"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/src/ship"
)

func main() {

	s := &ship.Ship{}

	s.rise(&ship.ShipCreated{})

	fmt.Printf("hola")
}
