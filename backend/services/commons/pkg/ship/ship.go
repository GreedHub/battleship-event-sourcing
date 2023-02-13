package ship

import "github.com/GreedHub/battleship-event-sourcing/backend/services/commons/pkg/utils"

type Ship struct {
	id       utils.ID
	name     string
	size     int
	position utils.Position
	hits     []utils.PosXY

	changes []ShipEvent
	version int
}

func New(id utils.ID, size int) *Ship {
	s := &Ship{}

	s.raise(&ShipCreated{
		ID:   id,
		Size: size,
	})

	return s
}

func (s *Ship) Position(position utils.Position) error {
	/* TODO: handle already positioned
	if s.position.Start != nil {
		return HandleAlreadyPositionedErr()
	}  */

	s.raise(&ShipPositioned{
		Position: position,
	})

	return nil
}

func (s *Ship) Hit(hit utils.PosXY) error {
	s.raise(&ShipHit{
		HitCoords: hit,
	})

	return nil
}

func (s *Ship) IsSink() bool {
	return len(s.hits) == s.size
}

func (s *Ship) IsHit() bool {
	return len(s.hits) > 0
}

func (s *Ship) Events() []ShipEvent {
	return s.changes
}

func (s *Ship) Version() int {
	return s.version
}
