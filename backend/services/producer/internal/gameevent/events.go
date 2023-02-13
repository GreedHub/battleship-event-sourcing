package gameevent

import (
	"errors"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/player"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/ship"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"
)

type GameEvent interface {
	isGameEvent()
	unmarshal(payload map[string]interface{})
}

func (e CreateSession)	isGameEvent() {}
func (e DeleteSession)	isGameEvent() {}
func (e JoinSession)	isGameEvent() {}
func (e ExitSession)	isGameEvent() {}
func (e StartMatch)		isGameEvent() {}
func (e PlaceShip)		isGameEvent() {}
func (e PlayerReady)	isGameEvent() {}
func (e PlayerShoot)	isGameEvent() {}

// CreateSession event.
type CreateSession struct {
	OwnerID player.PlayerID `json:"owner_id"`
}

func (e *CreateSession) unmarshal(payload map[string]interface{}) error{
	if payload["owner_id"] == nil{
		return errors.New("Error invoking envent CreateSession: Parameter 'owner_id' missing")
	}

	e.OwnerID = int(payload["owner_id"].(float64))

	return nil
}

// DeleteSession event.
type DeleteSession struct {
	OwnerID player.PlayerID `json:"owner_id"`
	SessionID  string `json:"session_id"`
}

func (e *DeleteSession) unmarshal(payload map[string]interface{}) error{
	if payload["owner_id"] == nil{
		return errors.New("Error invoking envent DeleteSession: Parameter 'owner_id' missing")
	}

	if payload["session_id"] == nil{
		return errors.New("Error invoking envent DeleteSession: Parameter 'session_id' missing")
	}

	e.OwnerID = int(payload["owner_id"].(float64))
	e.SessionID = payload["session_id"].(string)

	return nil
}

// JoinSession event.
type JoinSession struct {
	GuestID player.PlayerID `json:"guest_id"`
	SessionID  string `json:"session_id"`
}

func (e *JoinSession) unmarshal(payload map[string]interface{}) error{
	if payload["session_id"] == nil{
		return errors.New("Error invoking envent JoinSession: Parameter 'session_id' missing")
	}
	
	if payload["guest_id"] == nil{
		return errors.New("Error invoking envent JoinSession: Parameter 'guest_id' missing")
	}

	e.GuestID= int(payload["guest_id"].(float64))
	e.SessionID= payload["session_id"].(string)

	return nil
}

// ExitSession event.
type ExitSession struct {
	GuestID player.PlayerID `json:"guest_id"`
	SessionID  string `json:"session_id"`
}

func (e *ExitSession) unmarshal(payload map[string]interface{}) error{
	if payload["session_id"] == nil{
		return errors.New("Error invoking envent ExitSession: Parameter 'session_id' missing")
	}
	
	if payload["guest_id"] == nil{
		return errors.New("Error invoking envent ExitSession: Parameter 'guest_id' missing")
	}

	e.GuestID= int(payload["guest_id"].(float64))
	e.SessionID= payload["session_id"].(string)

	return nil
}

// StartMatch event.
type StartMatch struct {
	OwnerID player.PlayerID `json:"owner_id"`
	SessionID  string `json:"session_id"`
}

func (e *StartMatch) unmarshal(payload map[string]interface{}) error{
	if payload["owner_id"] == nil{
		return errors.New("Error invoking envent StartMatch: Parameter 'owner_id' missing")
	}

	if payload["session_id"] == nil{
		return errors.New("Error invoking envent StartMatch: Parameter 'session_id' missing")
	}

	e.OwnerID = int(payload["owner_id"].(float64))
	e.SessionID = payload["session_id"].(string)

	return nil
}

// PlaceShip event.
type PlaceShip struct {
	PlayerID player.PlayerID `json:"player_id"`
	SessionID  string `json:"session_id"`
	ShipID ship.ShipID `json:"ship_id"`
	Position utils.Position `json:"position"`
}

func (e *PlaceShip) unmarshal(payload map[string]interface{}) error{
	if payload["session_id"] == nil{
		return errors.New("Error invoking envent PlaceShip: Parameter 'session_id' missing")
	}
	
	if payload["player_id"] == nil{
		return errors.New("Error invoking envent PlaceShip: Parameter 'player_id' missing")
	}

	if payload["ship_id"] == nil{
		return errors.New("Error invoking envent PlaceShip: Parameter 'ship_id' missing")
	}

	if payload["position"] == nil{
		return errors.New("Error invoking envent PlaceShip: Parameter 'position' missing")
	}

	e.PlayerID	= int(payload["player_id"].(float64))
	e.ShipID	= int(payload["ship_id"].(float64))
	e.SessionID	= payload["session_id"].(string)
	e.Position	= payload["position"].(utils.Position)

	return nil
}

// PlayerReady event.
type PlayerReady struct {
	PlayerID player.PlayerID `json:"player_id"`
	SessionID  string `json:"session_id"`
}

func (e *PlayerReady) unmarshal(payload map[string]interface{}) error{
	if payload["player_id"] == nil{
		return errors.New("Error invoking envent PlayerReady: Parameter 'player_id' missing")
	}

	if payload["session_id"] == nil{
		return errors.New("Error invoking envent PlayerReady: Parameter 'session_id' missing")
	}

	e.PlayerID = int(payload["player_id"].(float64))
	e.SessionID = payload["session_id"].(string)

	return nil
}

// PlayerShoot event.
type PlayerShoot struct {
	PlayerID player.PlayerID `json:"player_id"`
	SessionID  string `json:"session_id"`
	Coords utils.PosXY `json:"coords"`
}

func (e *PlayerShoot) unmarshal(payload map[string]interface{}) error{
	if payload["player_id"] == nil{
		return errors.New("Error invoking envent PlayerShoot: Parameter 'player_id' missing")
	}

	if payload["session_id"] == nil{
		return errors.New("Error invoking envent PlayerShoot: Parameter 'session_id' missing")
	}

	if payload["coords"] == nil{
		return errors.New("Error invoking envent PlayerShoot: Parameter 'coords' missing")
	}

	e.PlayerID	= int(payload["player_id"].(float64))
	e.SessionID = payload["session_id"].(string)
	e.Coords	= payload["coords"].(utils.PosXY)

	return nil
}