package session

type SessionVerification int

const (
	NoVerify SessionVerification = iota
	AllowList
	DenyList
)

var sessionVerificationNames = [...]string{
	"NoVerify",
	"AllowList",
	"DenyList",
}
