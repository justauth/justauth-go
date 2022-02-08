package request

type AuthWxRequest struct {
	AuthDefaultRequest
}

func (d *AuthWxRequest) Authorize() string {
	return "Dota"
}
