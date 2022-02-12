package request

import (
	"fmt"
	"github.com/justauth/justauth-go/model"
)

type AuthDefaultRequest struct {
	authConfig model.AuthConfig
}

func (adq *AuthDefaultRequest) Authorize(state string) string {
	return fmt.Sprintf("%s")
}

func (adq *AuthDefaultRequest) Login(authCallBack model.AuthCallback) model.AuthResponse {
	authToken := adq.getAccessToken(authCallBack)
	userInfo := adq.getUserInfo(authToken)
	return model.AuthResponse{
		Code: 200,
		Msg:  "",
		Data: userInfo,
	}
}

func (adq *AuthDefaultRequest) Revoke(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (adq *AuthDefaultRequest) Refresh(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (adq *AuthDefaultRequest) getAccessToken(authCallback model.AuthCallback) model.AuthToken {
	return model.AuthToken{}
}

func (adq *AuthDefaultRequest) getUserInfo(authToken model.AuthToken) model.AuthUser {
	return model.AuthUser{}
}
