package ship

import "github.com/GreedHub/battleship-event-sourcing/backend/services/commons/src/utils"

// ShipEvent is a domain event marker.
type ShipEvent interface {
	isShipEvent()
}

func (e ShipCreated) isShipEvent()    {}
func (e ShipPositioned) isShipEvent() {}
func (e ShipHit) isShipEvent()        {}

// ShipCreated event.
type ShipCreated struct {
	ID   int `json:"id"`
	Size int `json:"size"`
}

// ShipPositioned event.
type ShipPositioned struct {
	ID       int            `json:"id"`
	Position utils.Position `json:"position"`
}

// ShipHit event.
type ShipHit struct {
	ID        int         `json:"id"`
	HitCoords utils.PosXY `json:"coords"`
}
