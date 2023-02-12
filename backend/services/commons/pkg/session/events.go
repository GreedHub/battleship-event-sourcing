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

/* {
	"id": "nonse",
	"timestamp": "2022-2-2T03:03:03Z-3:00",
	"type":"SessionCreated",
	"session_id": "8241yhbrkj",
	"owner_id": "kjdlvblh",
} 
*/

// GuestConnected event.
type GuestConnected struct {
	Guest player.PlayerID `json:"guest"`
}

// GuestDisconnected event.
type GuestDisconnected struct {}

/* {
	"id": "nonse",
	"timestamp": "2022-2-2T03:03:03Z-3:00",
	"type":"GuestConnected",
	"session_id": "8241yhbrkj",
	"guest_id": "kjdlvbl2h",
} 
*/

// MatchStarted event.
type MatchStarted struct{}

/* {
	"id": "nonse",
	"timestamp": "2022-2-2T03:03:03Z-3:00",
	"type":"MatchStarted",
	"session_id": "8241yhbrkj",
} 
*/

// GuestReady event.
type GuestReady struct{}

/* {
	"id": "nonse",
	"timestamp": "2022-2-2T03:03:03Z-3:00",
	"type":"GuestReady",
	"session_id": "8241yhbrkj",
} 
*/

// OwnerReady event.
type OwnerReady struct{}

// GuestWon event.
type GuestWon struct{}

// OwnerWon event.
type OwnerWon struct{}
