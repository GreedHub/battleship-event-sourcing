package ship

func (s *Ship) raise(event ShipEvent) {
	s.changes = append(s.changes, event)
	s.On(event, true)
}

func (s *Ship) On(event ShipEvent, new bool) {
	switch e := event.(type) {
	case *ShipCreated:
		s.onShipCreated(e)

	case *ShipPositioned:
		s.onShipPositioned(e)

	case *ShipHit:
		s.onShipHit(e)

	}

	if !new {
		s.version++
	}
}

func (s *Ship) onShipCreated(e *ShipCreated) {
	s.id = e.ID
	s.size = e.Size
}

func (s *Ship) onShipPositioned(e *ShipPositioned) {
	s.position = e.Position
}

func (s *Ship) onShipHit(e *ShipHit) {
	s.hits = append(s.hits, e.HitCoords)
}
