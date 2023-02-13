package gameevent

import (
	"errors"
	"fmt"
	"time"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/domain"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/player"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/session"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/producer/internal/kafka"
	"github.com/gin-gonic/gin"
)

type EventRequest struct {
	EventType string                 `json:"type"`
	Payload   map[string]interface{} `json:"payload"`
}

type EventProduction struct {
	Id        utils.ID    `json:"id"`
	Timestamp time.Time   `json:"timestamp"`
	EventType string      `json:"type"`
	EntityID  utils.ID    `json:"entity_id"`
	Payload   interface{} `json:"payload"`
}

func newEventProduction(eventType string, entityID utils.ID, event interface{}) *EventProduction {
	production := &EventProduction{
		Id:        utils.GetRandomString(50), // Fixme: this can produce duplicates
		Timestamp: time.Now(),
		EventType: eventType,
		EntityID:  entityID,
		Payload:   event,
	}
	return production
}

func HandleGameEvent(c *gin.Context) (status int, body map[string]interface{}, err error) {
	var ev EventRequest
	c.BindJSON(&ev)
	switch ev.EventType {
	case CREATE_PLAYER_EVENT:
		e := &CreatePlayer{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onCreatePlayerEvent(e)

	case CREATE_SESSION_EVENT:
		e := &CreateSession{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onCreateSessionEvent(e)

	case DELETE_SESSION_EVENT:

		e := &DeleteSession{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onDeleteSessionEvent(e)

	case JOIN_SESSION_EVENT:
		e := &JoinSession{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onJoinSessionEvent(e)

	case EXIT_SESSION_EVENT:
		e := &ExitSession{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onExitSessionEvent(e)

	case START_MATCH_EVENT:
		e := &StartMatch{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onStartMatchEvent(e)

	case PLACE_SHIP_EVENT:
		e := &PlaceShip{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onPlaceShipEvent(e)

	case PLAYER_READY_EVENT:
		e := &PlayerReady{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onPlayerReadyEvent(e)

		status, body, err = onPlayerReadyEvent(&PlayerReady{
			PlayerID:  ev.Payload["player_id"].(string),
			SessionID: ev.Payload["session_id"].(string),
		})

	case PLAYER_SHOOT_EVENT:
		e := &PlayerShoot{}
		err := e.unmarshal(ev.Payload)

		if err != nil {
			return 400, nil, err
		}

		status, body, err = onPlayerShootEvent(e)

	default:
		return 400, nil, errors.New(fmt.Sprintf(`Invalid event "%s"`, ev.EventType))
	}

	return status, body, err
}

func onCreatePlayerEvent(e *CreatePlayer) (status int, body map[string]interface{}, err error) {
	status = 201

	playerID := player.CreatePlayerId()
	body = make(map[string]interface{})
	body["player_id"] = playerID

	event := &player.PlayerCreated{
		Name:     e.Name,
		PlayerID: playerID,
	}

	production := newEventProduction("PlayerCreated", playerID, event)

	err = kafka.Produce(domain.PLAYER, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}

func onCreateSessionEvent(e *CreateSession) (status int, body map[string]interface{}, err error) {
	status = 201

	sessionId := session.CreateSessionId()
	body = make(map[string]interface{})
	body["session_id"] = sessionId

	event := &session.SessionCreated{
		OwnerID:   e.OwnerID,
		SessionID: sessionId,
	}

	production := newEventProduction("SessionCreated", sessionId, event)

	err = kafka.Produce(domain.SESSION, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}

func onDeleteSessionEvent(e *DeleteSession) (status int, body map[string]interface{}, err error) {
	status = 202

	event := &session.SessionDeleted{}

	production := newEventProduction("SessionDeleted", e.SessionID, event)

	err = kafka.Produce(domain.SESSION, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}

func onJoinSessionEvent(e *JoinSession) (status int, body map[string]interface{}, err error) {
	status = 200

	event := &session.GuestConnected{
		GuestID: e.GuestID,
	}

	production := newEventProduction("GuestConnected", e.SessionID, event)

	err = kafka.Produce(domain.SESSION, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}

func onExitSessionEvent(e *ExitSession) (status int, body map[string]interface{}, err error) {
	status = 202

	event := &session.GuestDisconnected{}

	production := newEventProduction("GuestDisconnected", e.SessionID, event)

	err = kafka.Produce(domain.SESSION, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}

func onStartMatchEvent(e *StartMatch) (status int, body map[string]interface{}, err error) {
	status = 202

	event := &session.MatchStarted{}

	production := newEventProduction("MatchStarted", e.SessionID, event)

	err = kafka.Produce(domain.SESSION, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}

func onPlaceShipEvent(e *PlaceShip) (status int, body map[string]interface{}, err error) {
	status = 202

	event := &player.ShipPlaced{
		PlayerID: e.PlayerID,
		ShipID:   e.ShipID,
		Position: e.Position,
	}

	production := newEventProduction("ShipPlaced", e.SessionID, event)

	err = kafka.Produce(domain.SESSION, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}

func onPlayerReadyEvent(e *PlayerReady) (status int, body map[string]interface{}, err error) {
	status = 202

	var event interface{}
	var production *EventProduction

	s := session.New(e.PlayerID, session.CreateSessionId()) // FIXME: change this for a session validation

	if e.PlayerID == s.GetOwnerID() {
		event = &session.OwnerReady{}
		production = newEventProduction("OwnerReady", e.SessionID, event)
	} else {
		event = &session.GuestReady{}
		production = newEventProduction("GuestReady", e.SessionID, event)
	}

	// FIXME: Validate game started

	err = kafka.Produce(domain.SESSION, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}

func onPlayerShootEvent(e *PlayerShoot) (status int, body map[string]interface{}, err error) {
	status = 202

	event := &player.PlayerShoot{
		PlayerID: e.PlayerID,
		Coords:   e.Coords,
	}

	// FIXME: Validate game finished

	production := newEventProduction("PlayerShoot", e.SessionID, event)

	err = kafka.Produce(domain.SESSION, production)

	if err != nil {
		return 500, nil, err
	}

	return status, body, err
}
