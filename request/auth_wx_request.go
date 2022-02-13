package request

type AuthWxRequest struct {
	BaseAuthRequest
}

func (d *AuthWxRequest) Authorize(state string) string {
	return "AuthWxRequest-Authorize"
}
