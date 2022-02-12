package request

type AuthWxRequest struct {
	AuthDefaultRequest
}

func (d *AuthWxRequest) Authorize(state string) string {
	return "AuthWxRequest-Authorize"
}
