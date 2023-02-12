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
		PlayerID:   id,
		Name: name,
	})

	return p
}

func (p *Player) GetName() string{
	return p.name
}