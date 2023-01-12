package gameevent

import "github.com/GreedHub/battleship-event-sourcing/backend/services/commons/src/domain"

type GameEvent struct{
	Entity string `json:"entity"`
	Event RequestEvent `json:"event"`
}

type RequestEvent struct{
	Type string `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

func HandleGameEvent (event GameEvent){
	switch event.Entity{
	case domain.PLAYER:
		HandlePlayerEvent(event.Event)

			case domain.SESSION:
		HandleSessionEvent(event.Event)

			case domain.SHIP:
		HandleShipEvent(event.Event)
	}
}