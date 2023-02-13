package gameevent

import (
	"errors"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"
)

type GameEvent interface {
	isGameEvent()
	unmarshal(payload map[string]interface{})
}

const (
	CREATE_PLAYER_EVENT  = "CreatePlayer"
	CREATE_SESSION_EVENT = "CreateSession"
	DELETE_SESSION_EVENT = "DeleteSession"
	JOIN_SESSION_EVENT   = "JoinSession"
	EXIT_SESSION_EVENT   = "ExitSession"
	START_MATCH_EVENT    = "StartMatch"
	PLACE_SHIP_EVENT     = "PlaceShip"
	PLAYER_READY_EVENT   = "PlayerReady"
	PLAYER_SHOOT_EVENT   = "PlayerShoot"
)

func (e CreatePlayer) isGameEvent()  {}
func (e CreateSession) isGameEvent() {}
func (e DeleteSession) isGameEvent() {}
func (e JoinSession) isGameEvent()   {}
func (e ExitSession) isGameEvent()   {}
func (e StartMatch) isGameEvent()    {}
func (e PlaceShip) isGameEvent()     {}
func (e PlayerReady) isGameEvent()   {}
func (e PlayerShoot) isGameEvent()   {}

// CreatePlayer event.
type CreatePlayer struct {
	Name string `json:"name"`
}

func (e *CreatePlayer) unmarshal(payload map[string]interface{}) error {
	if payload["name"] == nil {
		return errors.New("Error invoking envent CreatePlayer: Parameter 'name' missing")
	}

	e.Name = payload["name"].(string)

	return nil
}

// CreateSession event.
type CreateSession struct {
	OwnerID utils.ID `json:"owner_id"`
}

func (e *CreateSession) unmarshal(payload map[string]interface{}) error {
	if payload["owner_id"] == nil {
		return errors.New("Error invoking envent CreateSession: Parameter 'owner_id' missing")
	}

	e.OwnerID = payload["owner_id"].(string)

	return nil
}

// DeleteSession event.
type DeleteSession struct {
	OwnerID   utils.ID `json:"owner_id"`
	SessionID utils.ID `json:"session_id"`
}

func (e *DeleteSession) unmarshal(payload map[string]interface{}) error {
	if payload["owner_id"] == nil {
		return errors.New("Error invoking envent DeleteSession: Parameter 'owner_id' missing")
	}

	if payload["session_id"] == nil {
		return errors.New("Error invoking envent DeleteSession: Parameter 'session_id' missing")
	}

	e.OwnerID = payload["owner_id"].(string)
	e.SessionID = payload["session_id"].(string)

	return nil
}

// JoinSession event.
type JoinSession struct {
	GuestID   utils.ID `json:"guest_id"`
	SessionID utils.ID `json:"session_id"`
}

func (e *JoinSession) unmarshal(payload map[string]interface{}) error {
	if payload["session_id"] == nil {
		return errors.New("Error invoking envent JoinSession: Parameter 'session_id' missing")
	}

	if payload["guest_id"] == nil {
		return errors.New("Error invoking envent JoinSession: Parameter 'guest_id' missing")
	}

	e.GuestID = payload["guest_id"].(string)
	e.SessionID = payload["session_id"].(string)

	return nil
}

// ExitSession event.
type ExitSession struct {
	GuestID   utils.ID `json:"guest_id"`
	SessionID utils.ID `json:"session_id"`
}

func (e *ExitSession) unmarshal(payload map[string]interface{}) error {
	if payload["session_id"] == nil {
		return errors.New("Error invoking envent ExitSession: Parameter 'session_id' missing")
	}

	if payload["guest_id"] == nil {
		return errors.New("Error invoking envent ExitSession: Parameter 'guest_id' missing")
	}

	e.GuestID = payload["guest_id"].(string)
	e.SessionID = payload["session_id"].(string)

	return nil
}

// StartMatch event.
type StartMatch struct {
	OwnerID   utils.ID `json:"owner_id"`
	SessionID utils.ID `json:"session_id"`
}

func (e *StartMatch) unmarshal(payload map[string]interface{}) error {
	if payload["owner_id"] == nil {
		return errors.New("Error invoking envent StartMatch: Parameter 'owner_id' missing")
	}

	if payload["session_id"] == nil {
		return errors.New("Error invoking envent StartMatch: Parameter 'session_id' missing")
	}

	e.OwnerID = payload["owner_id"].(string)
	e.SessionID = payload["session_id"].(string)

	return nil
}

// PlaceShip event.
type PlaceShip struct {
	PlayerID  utils.ID       `json:"player_id"`
	SessionID utils.ID       `json:"session_id"`
	ShipID    utils.ID       `json:"ship_id"`
	Position  utils.Position `json:"position"`
}

func (e *PlaceShip) unmarshal(payload map[string]interface{}) error {
	if payload["session_id"] == nil {
		return errors.New("Error invoking envent PlaceShip: Parameter 'session_id' missing")
	}

	if payload["player_id"] == nil {
		return errors.New("Error invoking envent PlaceShip: Parameter 'player_id' missing")
	}

	if payload["ship_id"] == nil {
		return errors.New("Error invoking envent PlaceShip: Parameter 'ship_id' missing")
	}

	if payload["position"] == nil {
		return errors.New("Error invoking envent PlaceShip: Parameter 'position' missing")
	}

	e.PlayerID = payload["player_id"].(string)
	e.ShipID = payload["ship_id"].(string)
	e.SessionID = payload["session_id"].(string)
	e.Position = payload["position"].(utils.Position)

	return nil
}

// PlayerReady event.
type PlayerReady struct {
	PlayerID  utils.ID `json:"player_id"`
	SessionID utils.ID `json:"session_id"`
}

func (e *PlayerReady) unmarshal(payload map[string]interface{}) error {
	if payload["player_id"] == nil {
		return errors.New("Error invoking envent PlayerReady: Parameter 'player_id' missing")
	}

	if payload["session_id"] == nil {
		return errors.New("Error invoking envent PlayerReady: Parameter 'session_id' missing")
	}

	e.PlayerID = payload["player_id"].(string)
	e.SessionID = payload["session_id"].(string)

	return nil
}

// PlayerShoot event.
type PlayerShoot struct {
	PlayerID  utils.ID    `json:"player_id"`
	SessionID utils.ID    `json:"session_id"`
	Coords    utils.PosXY `json:"coords"`
}

func (e *PlayerShoot) unmarshal(payload map[string]interface{}) error {
	if payload["player_id"] == nil {
		return errors.New("Error invoking envent PlayerShoot: Parameter 'player_id' missing")
	}

	if payload["session_id"] == nil {
		return errors.New("Error invoking envent PlayerShoot: Parameter 'session_id' missing")
	}

	if payload["coords"] == nil {
		return errors.New("Error invoking envent PlayerShoot: Parameter 'coords' missing")
	}

	e.PlayerID = payload["player_id"].(string)
	e.SessionID = payload["session_id"].(string)
	e.Coords = payload["coords"].(utils.PosXY)

	return nil
}
