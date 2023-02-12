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

type EventRequest struct{
	EventType 	string		`json:"type"`
	Payload		map[string]interface{}	`json:"payload"`
}

type EventProduction struct{
	Id			eventID				`json:"id"`
	Timestamp	time.Time			`json:"timestamp"`
	EventType	string				`json:"type"`
	SessionId	session.SessionID	`json:"session_id"`
	Payload		interface{}			`json:"payload"`
}

type eventID = string

func newEventProduction(eventType string, sessionId string, event interface{}) *EventProduction{
	production := &EventProduction{
			Id: utils.GetRandomString(50), // Fixme: this can produce duplicates
			Timestamp: time.Now(),
			EventType: eventType,
			SessionId: sessionId,
			Payload: event,
		}
	return production
}

func HandleGameEvent (c *gin.Context)  (status int, body map[string]interface{}, err error){
	var ev EventRequest
	c.BindJSON(&ev)
	switch  ev.EventType{
	case "CreateSession":
		if ev.Payload["owner_id"] == nil{
			return 400,nil,errors.New("Error invoking envent CreateSession: Parameter 'owner_id' missing")
		}
		status, body, err = onCreateSessionEvent(&CreateSession{
			OwnerID: int(ev.Payload["owner_id"].(float64)),
		})

	case "DeleteSession":
		if ev.Payload["owner_id"] == nil{
			return 400,nil,errors.New("Error invoking envent DeleteSession: Parameter 'owner_id' missing")
		}
		if ev.Payload["session_id"] == nil{
			return 400,nil,errors.New("Error invoking envent JoinSession: Parameter 'session_id' missing")
		}

		status, body, err = onDeleteSessionEvent(&DeleteSession{
			OwnerID: int(ev.Payload["owner_id"].(float64)),
			SessionID: ev.Payload["session_id"].(string),
		})

	case "JoinSession":
		if ev.Payload["session_id"] == nil{
			return 400,nil,errors.New("Error invoking envent JoinSession: Parameter 'session_id' missing")
		}
		
		if ev.Payload["guest_id"] == nil{
			return 400,nil,errors.New("Error invoking envent JoinSession: Parameter 'guest_id' missing")
		}

		status, body, err = onJoinSessionEvent(&JoinSession{
			GuestID: int(ev.Payload["guest_id"].(float64)),
			SessionID: ev.Payload["session_id"].(string),
		})

	case "ExitSession":
		if ev.Payload["session_id"] == nil{
			return 400,nil,errors.New("Error invoking envent ExitSession: Parameter 'session_id' missing")
		}
		if ev.Payload["guest_id"] == nil{
			return 400,nil,errors.New("Error invoking envent ExitSession: Parameter 'guest_id' missing")
		}
		status, body, err = onExitSessionEvent(&ExitSession{
			GuestID: int(ev.Payload["guest_id"].(float64)),
			SessionID: ev.Payload["session_id"].(string),
		})

	case "StartMatch":
		if ev.Payload["session_id"] == nil{
			return 400,nil,errors.New("Error invoking envent StartMatch: Parameter 'session_id' missing")
		}
		
		if ev.Payload["owner_id"] == nil{
			return 400,nil,errors.New("Error invoking envent StartMatch: Parameter 'owner_id' missing")
		}

		status, body, err = onStartMatchEvent(&StartMatch{
			OwnerID: int(ev.Payload["owner_id"].(float64)),
			SessionID: ev.Payload["session_id"].(string),
		})

	case "PlaceShip":
		if ev.Payload["session_id"] == nil{
			return 400,nil,errors.New("Error invoking envent PlaceShip: Parameter 'session_id' missing")
		}
		
		if ev.Payload["player_id"] == nil{
			return 400,nil,errors.New("Error invoking envent PlaceShip: Parameter 'player_id' missing")
		}

		if ev.Payload["ship_id"] == nil{
			return 400,nil,errors.New("Error invoking envent PlaceShip: Parameter 'ship_id' missing")
		}

		if ev.Payload["position"] == nil{
			return 400,nil,errors.New("Error invoking envent PlaceShip: Parameter 'position' missing")
		}

		status, body, err = onPlaceShipEvent(&PlaceShip{
			PlayerID: int(ev.Payload["player_id"].(float64)),
			ShipID: int(ev.Payload["ship_id"].(float64)),
			SessionID: ev.Payload["session_id"].(string),
			Position: ev.Payload["position"].(utils.Position),
		})

	case "PlayerReady":
		if ev.Payload["session_id"] == nil{
			return 400,nil,errors.New("Error invoking envent PlaceShip: Parameter 'session_id' missing")
		}
		
		if ev.Payload["player_id"] == nil{
			return 400,nil,errors.New("Error invoking envent PlaceShip: Parameter 'player_id' missing")
		}

		status, body, err = onPlayerReadyEvent(&PlayerReady{
			PlayerID: int(ev.Payload["player_id"].(float64)),
			SessionID: ev.Payload["session_id"].(string),
		})

	case "PlayerShoot":
		if ev.Payload["session_id"] == nil{
			return 400,nil,errors.New("Error invoking envent PlayerShoot: Parameter 'session_id' missing")
		}
		
		if ev.Payload["player_id"] == nil{
			return 400,nil,errors.New("Error invoking envent PlayerShoot: Parameter 'player_id' missing")
		}
		
		if ev.Payload["coords"] == nil{
			return 400,nil,errors.New("Error invoking envent PlayerShoot: Parameter 'coords' missing")
		}

		status, body, err = onPlayerShootEvent(&PlayerShoot{
			PlayerID: int(ev.Payload["player_id"].(float64)),
			SessionID: ev.Payload["session_id"].(string),
			Coords: ev.Payload["coords"].(utils.PosXY),
		})

	default:
		return 400, nil, errors.New(fmt.Sprintf("Invalid event: %s",ev.EventType)) 
	}

	return status, body, err
}

func onCreateSessionEvent(e *CreateSession) (status int, body map[string]interface{}, err error){
	status = 201

	sessionId  := session.CreateSessionId()
	body = make(map[string]interface{})
	body["session_id"] = sessionId

	event := &session.SessionCreated{
		OwnerID: e.OwnerID,
		SessionID: sessionId,
	}
	
	production := newEventProduction("SessionCreated",sessionId,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onDeleteSessionEvent(e *DeleteSession)(status int, body map[string]interface{}, err error){
	status = 202

	event := &session.SessionDeleted{}
	
	production := newEventProduction("SessionDeleted",e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onJoinSessionEvent(e *JoinSession)(status int, body map[string]interface{}, err error){
	status = 200

	event := &session.GuestConnected{
		GuestID: e.GuestID,
	}
	
	production := newEventProduction("GuestConnected",e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onExitSessionEvent(e *ExitSession)(status int, body map[string]interface{}, err error){
	status = 202

	event := &session.GuestDisconnected{}
	
	production := newEventProduction("GuestDisconnected",e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onStartMatchEvent(e *StartMatch)(status int, body map[string]interface{}, err error){
	status = 202

	event := &session.MatchStarted{}
	
	production := newEventProduction("MatchStarted",e.SessionID,event)

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
	
	production := newEventProduction("ShipPlaced",e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}

func onPlayerReadyEvent(e *PlayerReady)(status int, body map[string]interface{}, err error){
	status = 202

	var event interface{}
	var production *EventProduction

	s := session.New(e.PlayerID,session.CreateSessionId()) // FIXME: change this for a session validation

	if e.PlayerID == s.GetOwnerID(){
		event = &session.OwnerReady{}
		production = newEventProduction("OwnerReady",e.SessionID,event)
	} else {
		event = &session.GuestReady{}
		production = newEventProduction("GuestReady",e.SessionID,event)
	}

	// FIXME: Validate game started

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
	
	production := newEventProduction("PlayerShoot",e.SessionID,event)

	err = kafka.Produce(domain.SESSION,production)

	if err != nil { 
		return 500, nil, err
	}
	
	return status, body, err
}