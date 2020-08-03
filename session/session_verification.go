package session

type SessionVerification int

const (
	NoVerify SessionVerification = iota
	WhiteList
	BlackList
)

var sessionVerificationNames = [...]string{
	"NoVerify",
	"WhiteList",
	"BlackList",
}
