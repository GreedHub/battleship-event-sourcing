package player

// PlayerEvent is a domain event marker.
type PlayerEvent interface {
	isPlayerEvent()
}

func (e PlayerCreated) isPlayerEvent() {}

// PlayerCreated event.
type PlayerCreated struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
