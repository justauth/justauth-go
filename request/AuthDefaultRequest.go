package request

import "fmt"

type AuthDefaultRequest struct{}

func (a *AuthDefaultRequest) Play(authRequest AuthRequest) {
	fmt.Printf(fmt.Sprintf("%s is awesome!", authRequest.Authorize()))
}
