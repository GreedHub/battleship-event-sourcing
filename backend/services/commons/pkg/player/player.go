package player

import (
	"os"
	"strconv"

	"github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"
)

/*

 */
type Player struct {
	id   PlayerID
	name string

	changes []PlayerEvent
	version int
}

type PlayerID = string

func New(id PlayerID, name string) *Player {
	p := &Player{}

	p.raise(&PlayerCreated{
		PlayerID:   id,
		Name: name,
	})

	return p
}

func CreatePlayerId() PlayerID {
	DEFAULT_PLAYER_ID_LENGTH := 50

	playerIdLengthEnv := os.Getenv("Player_ID_LENGTH")
	playerIdLength, err := strconv.Atoi(playerIdLengthEnv)

	if err != nil {
		playerIdLength = DEFAULT_PLAYER_ID_LENGTH
	}

	if playerIdLength < DEFAULT_PLAYER_ID_LENGTH {
		playerIdLength = DEFAULT_PLAYER_ID_LENGTH
	}

	return utils.GetRandomString(playerIdLength)
}

func (p *Player) GetName() string{
	return p.name
}