package gameevent

import (
	"errors"
	"reflect"
	"time"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/domain"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/player"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/session"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/producer/internal/kafka"
)

type EventProduction struct{
	Id			eventID				`json:"id"`
	Timestamp	time.Time			`json:"timestamp"`
	EventType	string				`json:"type"`
	SessionId	session.SessionID	`json:"session_id"`
	Payload		interface{}			`json:"payload"`
}

type eventID = string

func newEventProduction(sessionId string, event interface{}) *EventProduction{
	production := &EventProduction{
			Id: utils.GetRandomString(50), // Fixme: this can produce duplicates
			Timestamp: time.Now(),
			EventType: reflect.TypeOf(event).Name(),
			SessionId: sessionId,
			Payload: event,
		}
	return production
}


func HandleGameEvent (gameEvent interface{})  (status int, body map[string]interface{}, err error){
	switch  e:= gameEvent.(type){
	case *CreateSession:
		status, body, err = onCreateSessionEvent(e)

	case *DeleteSession:
		status, body, err = onDeleteSessionEvent(e)

	case *JoinSession:
		status, body, err = onJoinSessionEvent(e)

	case *ExitSession:
		status, body, err = onExitSessionEvent(e)

	case *StartMatch:
		status, body, err = onStartMatchEvent(e)

	case *PlaceShip:
		status, body, err = onPlaceShipEvent(e)

	case *PlayerReady:
		status, body, err = onPlayerReadyEvent(e)

	case *PlayerShoot:
		status, body, err = onPlayerShootEvent(e)

	default:
		return 400, nil, errors.New("Invalid event") 
	}

	return status, body, err
}

func onCreateSessionEvent(e *CreateSession) (status int, body map[string]interface{}, err error){
	status = 201

	sessionId  := session.CreateSessionId()
	body["session_id"] = sessionId

	event := &session.SessionCreated{
		OwnerID: e.OwnerID,
		SessionID: e.SessionID,
	}
	
	production := newEventProduction(sessionId,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onDeleteSessionEvent(e *DeleteSession)(status int, body map[string]interface{}, err error){
	status = 202

	event := &session.SessionDeleted{}
	
	production := newEventProduction(e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onJoinSessionEvent(e *JoinSession)(status int, body map[string]interface{}, err error){
	status = 200

	event := &session.SessionDeleted{}
	
	production := newEventProduction(e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onExitSessionEvent(e *ExitSession)(status int, body map[string]interface{}, err error){
	status = 202

	event := &session.GuestDisconnected{}
	
	production := newEventProduction(e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onStartMatchEvent(e *StartMatch)(status int, body map[string]interface{}, err error){
	status = 202

	event := &session.MatchStarted{}
	
	production := newEventProduction(e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}
func onPlaceShipEvent(e *PlaceShip)(status int, body map[string]interface{}, err error){
	status = 202

	event := &player.ShipPlaced{
		PlayerID: e.PlayerID,
		ShipID: e.ShipID,
		Position: e.Position,
	}
	
	production := newEventProduction(e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onPlayerReadyEvent(e *PlayerReady)(status int, body map[string]interface{}, err error){
	status = 202

	var event interface{}

	s := session.New(e.PlayerID,session.CreateSessionId()) // FIXME: change this for a session validation

	if e.PlayerID == s.GetOwnerID(){
		event = &session.OwnerReady{}
	} else {
		event = &session.GuestReady{}
	}

	// FIXME: Validate game started

	production := newEventProduction(e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onPlayerShootEvent(e *PlayerShoot)(status int, body map[string]interface{}, err error){
	status = 202

	event := &player.PlayerShoot{
		PlayerID: e.PlayerID,
		Coords: e.Coords,
	}

	// FIXME: Validate game finished
	
	production := newEventProduction(e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}