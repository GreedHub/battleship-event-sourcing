package session

import "github.com/GreedHub/battleship-event-sourcing/backend/services/commons/src/player"


type SessionEvent interface {
	isSessionEvent()
}

func (e SessionCreated) isSessionEvent() {}
func (e GuestConnected) isSessionEvent() {}
func (e MatchStarted) isSessionEvent()   {}
func (e GuestReady) isSessionEvent()     {}
func (e OwnerReady) isSessionEvent()     {}
func (e GuestWon) isSessionEvent()       {}
func (e OwnerWon) isSessionEvent()       {}

// SessionCreated event.
type SessionCreated struct {
	Owner player.PlayerID `json:"owner"`
}

// GuestConnected event.
type GuestConnected struct {
	Guest player.PlayerID `json:"guest"`
}

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
