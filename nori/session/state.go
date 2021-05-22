package session

type State uint8

const (
	SessionUndefined State = iota
	SessionClosed
	SessionActive
	SessionLocked
	SessionBlocked
	SessionExpired
	SessionError
)

var states = []string{"undefined", "active", "closed", "locked", "blocked", "expired", "error"}

func NewState(v uint8) State {
	if len(states) < int(v) {
		return SessionUndefined
	}
	return State(v)
}

func NewStateFromString(v string) State {
	for i := range states {
		if states[i] == v {
			return State(i)
		}
	}
	return SessionUndefined
}

func (s State) String() string {
	if len(states) < int(s) {
		return states[0]
	}
	return states[s]
}

func (s State) Value() uint8 {
	return uint8(s)
}
