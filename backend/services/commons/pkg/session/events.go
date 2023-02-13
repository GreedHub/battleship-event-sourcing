package session

import "github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/player"

type SessionEvent interface {
	isSessionEvent()
}

func (e SessionCreated) isSessionEvent() {}
func (e SessionDeleted) isSessionEvent() {}
func (e GuestConnected) isSessionEvent() {}
func (e GuestDisconnected) isSessionEvent() {}
func (e MatchStarted) isSessionEvent()   {}
func (e GuestReady) isSessionEvent()     {}
func (e OwnerReady) isSessionEvent()     {}
func (e GuestWon) isSessionEvent()       {}
func (e OwnerWon) isSessionEvent()       {}

// SessionCreated event.
type SessionCreated struct {
	OwnerID player.PlayerID `json:"owner"`
	SessionID  string `json:"session_id"`
}

// SessionDeleted event.
type SessionDeleted struct {}

// GuestConnected event.
type GuestConnected struct {
	GuestID player.PlayerID `json:"guest"`
}

// GuestDisconnected event.
type GuestDisconnected struct {}

// MatchStarted event.
type MatchStarted struct{}

// GuestReady event.
type GuestReady struct{}

// OwnerReady event.
type OwnerReady struct{}

// GuestWon event.
type GuestWon struct{}

// OwnerWon event.
type OwnerWon struct{}
