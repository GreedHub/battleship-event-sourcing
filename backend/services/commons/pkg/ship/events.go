package ship

import "github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"

// ShipEvent is a domain event marker.
type ShipEvent interface {
	isShipEvent()
}

func (e ShipCreated) isShipEvent()    {}
func (e ShipPositioned) isShipEvent() {}
func (e ShipHit) isShipEvent()        {}

// ShipCreated event.
type ShipCreated struct {
	ID   utils.ID `json:"id"`
	Size int `json:"size"`
}

// ShipPositioned event.
type ShipPositioned struct {
	ID       utils.ID            `json:"id"`
	Position utils.Position `json:"position"`
}

// ShipHit event.
type ShipHit struct {
	ID        utils.ID         `json:"id"`
	HitCoords utils.PosXY `json:"coords"`
}
