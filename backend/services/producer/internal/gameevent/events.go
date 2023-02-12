package gameevent

import (
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/player"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/ship"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"
)

type GameEvent interface {
	isGameEvent()
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

// DeleteSession event.
type DeleteSession struct {
	OwnerID player.PlayerID `json:"owner_id"`
	SessionID  string `json:"session_id"`
}

// JoinSession event.
type JoinSession struct {
	GuestID player.PlayerID `json:"guest_id"`
	SessionID  string `json:"session_id"`
}

// ExitSession event.
type ExitSession struct {
	GuestID player.PlayerID `json:"guest_id"`
	SessionID  string `json:"session_id"`
}

// StartMatch event.
type StartMatch struct {
	OwnerID player.PlayerID `json:"owner_id"`
	SessionID  string `json:"session_id"`
}

// PlaceShip event.
type PlaceShip struct {
	PlayerID player.PlayerID `json:"player_id"`
	SessionID  string `json:"session_id"`
	ShipID ship.ShipID `json:"ship_id"`
	Position utils.Position `json:"position"`
}

// PlayerReady event.
type PlayerReady struct {
	PlayerID player.PlayerID `json:"player_id"`
	SessionID  string `json:"session_id"`
}

// PlayerShoot event.
type PlayerShoot struct {
	PlayerID player.PlayerID `json:"player_id"`
	SessionID  string `json:"session_id"`
	Coords utils.PosXY `json:"coords"`
}