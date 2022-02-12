package request

import "github.com/justauth/justauth-go/model"

type AuthRequest interface {
	Authorize(state string) string
	Login(authCallBack model.AuthCallback) model.AuthResponse
	Revoke(authToken model.AuthToken) model.AuthResponse
	Refresh(authToken model.AuthToken) model.AuthResponse
}
