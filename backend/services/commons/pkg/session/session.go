package session

import (
	"os"
	"strconv"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/player"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"
)

type Session struct {
	id           SessionID
	ownerID        player.PlayerID
	guestID        player.PlayerID
	status       string
	playersReady int

	changes []SessionEvent
	version int
}

type SessionID = string

func New(ownerID player.PlayerID, sessionID SessionID) *Session {
	s := &Session{}

	s.raise(&SessionCreated{
		OwnerID: ownerID,
		SessionID: sessionID,
	})

	return s
}

func CreateSessionId() SessionID {
	DEFAULT_SESSION_ID_LENGTH := 4

	sessionIdLengthEnv := os.Getenv("SESSION_ID_LENGTH")
	sessionIdLength, err := strconv.Atoi(sessionIdLengthEnv)

	if err != nil {
		sessionIdLength = DEFAULT_SESSION_ID_LENGTH
	}

	if sessionIdLength < 4 {
		sessionIdLength = DEFAULT_SESSION_ID_LENGTH
	}

	return utils.GetRandomCapitalizedString(sessionIdLength)
}

func (s *Session) GetOwnerID() player.PlayerID {
	return s.ownerID
}

func (s *Session) Events() []SessionEvent {
	return s.changes
}

func (s *Session) Version() int {
	return s.version
}

func (s *Session) BothPlayersReady() bool {
	return s.playersReady == 2
}