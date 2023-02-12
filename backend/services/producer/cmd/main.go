package main

import (
	"fmt"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/player"
)

func main() {

	p := player.New(123,"pepe")

	fmt.Printf("hola %s", p.GetName())
}
