package player

func (p *Player) raise(event PlayerEvent) {
	p.changes = append(p.changes, event)
	p.On(event, true)
}

// TODO crear estos eventos en el eventos go del player
func (p *Player) On(event PlayerEvent, new bool) {
	switch e := event.(type) {
	case *PlayerCreated:
		p.onPlayerCreated(e)

	}

	if !new {
		p.version++
	}
}

func (p *Player) onPlayerCreated(e *PlayerCreated) {
	p.id = e.ID
	p.name = e.Name
}
