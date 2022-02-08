package request

type AuthRequest interface {
	Authorize() string
}
