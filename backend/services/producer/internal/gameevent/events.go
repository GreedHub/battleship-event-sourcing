package gameevent

import (
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
	requiredParams := []string{"name"}

	err := utils.ValidateRequestParams(CREATE_PLAYER_EVENT, requiredParams, payload)

	if err != nil {
		return err
	}

	e.Name = payload["name"].(string)

	return nil
}

// CreateSession event.
type CreateSession struct {
	OwnerID utils.ID `json:"owner_id"`
}

func (e *CreateSession) unmarshal(payload map[string]interface{}) error {
	requiredParams := []string{"owner_id"}

	err := utils.ValidateRequestParams(CREATE_SESSION_EVENT, requiredParams, payload)

	if err != nil {
		return err
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
	requiredParams := []string{"session_id", "owner_id"}

	err := utils.ValidateRequestParams(DELETE_SESSION_EVENT, requiredParams, payload)

	if err != nil {
		return err
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
	requiredParams := []string{"session_id", "guest_id"}

	err := utils.ValidateRequestParams(JOIN_SESSION_EVENT, requiredParams, payload)

	if err != nil {
		return err
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
	requiredParams := []string{"session_id", "guest_id"}

	err := utils.ValidateRequestParams(EXIT_SESSION_EVENT, requiredParams, payload)

	if err != nil {
		return err
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
	requiredParams := []string{"session_id", "owner_id"}

	err := utils.ValidateRequestParams(START_MATCH_EVENT, requiredParams, payload)

	if err != nil {
		return err
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
	requiredParams := []string{"session_id", "player_id", "ship_id", "position"}

	err := utils.ValidateRequestParams(PLACE_SHIP_EVENT, requiredParams, payload)

	if err != nil {
		return err
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

	requiredParams := []string{"session_id", "player_id"}

	err := utils.ValidateRequestParams(PLAYER_READY_EVENT, requiredParams, payload)

	if err != nil {
		return err
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
	requiredParams := []string{"session_id", "player_id", "coords"}

	err := utils.ValidateRequestParams(PLAYER_SHOOT_EVENT, requiredParams, payload)

	if err != nil {
		return err
	}

	e.PlayerID = payload["player_id"].(string)
	e.SessionID = payload["session_id"].(string)
	e.Coords = payload["coords"].(utils.PosXY)

	return nil
}
