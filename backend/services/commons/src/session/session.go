package session

import (
	"os"
	"strconv"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/src/player"
	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/src/utils"
)

type Session struct {
	id           SessionID
	owner        player.PlayerID
	guest        player.PlayerID
	status       string
	playersReady int

	changes []SessionEvent
	version int
}

type SessionID = string

func New(owner player.PlayerID) *Session {
	s := &Session{}

	s.raise(&SessionCreated{
		Owner: owner,
	})

	return s
}

func (s *Session) createSessionId() SessionID {
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

func (s *Session) Events() []SessionEvent {
	return s.changes
}

func (s *Session) Version() int {
	return s.version
}
