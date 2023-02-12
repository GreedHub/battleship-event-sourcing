package player

import (
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/ship"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"
)

// PlayerEvent is a domain event marker.
type PlayerEvent interface {
	isPlayerEvent()
}

func (e PlayerCreated) isPlayerEvent() {}
func (e ShipPlaced) isPlayerEvent() {}
func (e PlayerShoot) isPlayerEvent() {}


// PlayerCreated event.
type PlayerCreated struct {
	PlayerID   	PlayerID	`json:"id"`
	Name 		string		`json:"name"`
}

// ShipPlaced event.
type ShipPlaced struct {
	PlayerID   PlayerID			`json:"player_id"`
	ShipID	   ship.ShipID		`json:"ship_id"`
	Position   utils.Position	`json:"position"`
}

// PlayerShoot event.
type PlayerShoot struct {
	PlayerID	PlayerID	`json:"player_id"`
	Coords		utils.PosXY	`json:"coords"`
}