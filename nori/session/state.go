package session

type State uint8

const (
	SessionClosed State = iota
	SessionActive
	SessionLocked
	SessionBlocked
	SessionExpired
	SessionError
)

var states = []string{"active", "closed", "locked", "blocked", "expired", "error"}

func (s State) String() string {
	if len(states) < int(s) {
		return ""
	}
	return states[s]
}
