package session

type State string

const (
	SessionActive  State = "active"
	SessionClosed  State = "closed"
	SessionLocked  State = "locked"
	SessionBlocked State = "blocked"
	SessionExpired State = "expired"
	SessionError   State = "error"
)