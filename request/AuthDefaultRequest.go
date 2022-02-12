package request

import (
	"github.com/justauth/justauth-go/model"
)

type AuthDefaultRequest struct{}

func (d *AuthDefaultRequest) Authorize(state string) string {
	return "Authorize"
}

func (d *AuthDefaultRequest) Login(authCallBack model.AuthCallback) model.AuthResponse {
	return model.AuthResponse{}
}

func (d *AuthDefaultRequest) Revoke(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (d *AuthDefaultRequest) Refresh(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}
