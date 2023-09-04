package kafka

import (
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/domain"
)

var (
	playerEvents  []Production
	sessionEvents []Production
	shipEvents    []Production
)

type Production = interface{}

func Produce(queue string, event interface{}) error {
	switch queue {
	case domain.PLAYER:
		playerEvents = append(playerEvents, event)

	case domain.SESSION:
		sessionEvents = append(sessionEvents, event)

	case domain.SHIP:
		shipEvents = append(shipEvents, event)

	}

	return nil
}

func Read() (player []Production, session []Production, ship []Production) {
	return playerEvents, sessionEvents, shipEvents
}
