package session

func (s *Session) raise(event SessionEvent) {
	s.changes = append(s.changes, event)
	s.On(event, true)
}

func (s *Session) On(event SessionEvent, new bool) {
	switch e := event.(type) {
	case *SessionCreated:
		s.onSessionCreated(e)

	case *GuestConnected:
		s.onGuestConnected(e)

	case *MatchStarted:
		s.onMatchStarted(e)

	case *GuestReady:
		s.onGuestReady(e)

	case *OwnerReady:
		s.onOwnerReady(e)

	case *GuestWon:
		s.onGuestWon(e)

	case *OwnerWon:
		s.onOwnerWon(e)
	}

	if !new {
		s.version++
	}
}

func (s *Session) onSessionCreated(e *SessionCreated) {
	s.id = e.SessionID
	s.ownerID = e.OwnerID
	s.status = WAITING_FOR_GUEST
}

func (s *Session) onGuestConnected(e *GuestConnected) {
	s.guestID = e.GuestID
	s.status = WAITING_TO_START
}

func (s *Session) onMatchStarted(e *MatchStarted) {
	s.playersReady = 0
	s.status = WAITING_FOR_PLACEMENTS
}

func (s *Session) onGuestReady(e *GuestReady) {
	s.playersReady++

	if s.BothPlayersReady() {
		s.status = IN_GAME
		return
	}

	s.status = WAITING_FOR_OWNER_PLACEMENTS
}

func (s *Session) onOwnerReady(e *OwnerReady) {
	s.playersReady++

	if s.BothPlayersReady() {
		s.status = IN_GAME
		return
	}

	s.status = WAITING_FOR_GUEST_PLACEMENTS
}

func (s *Session) onGuestWon(e *GuestWon) {
	if s.status != IN_GAME {
		return // TODO: throw an error
	}

	s.status = FINISHED_GUEST_WON
}

func (s *Session) onOwnerWon(e *OwnerWon) {
	if s.status != IN_GAME {
		return // TODO: throw an error
	}

	s.status = FINISHED_OWNER_WON
}
