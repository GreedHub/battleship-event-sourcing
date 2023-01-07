package player

/*

 */
type Player struct {
	id   PlayerID
	name string

	changes []PlayerEvent
	version int
}

type PlayerID = int

func New(id int, name string) *Player {
	p := &Player{}

	p.raise(&PlayerCreated{
		ID:   id,
		Name: name,
	})

	return p
}
