package request

import (
	"fmt"
	"github.com/justauth/justauth-go/model"
)

type AuthGiteeRequest struct {
	adq AuthDefaultRequest
}

func (agq *AuthGiteeRequest) Authorize(state string) string {
	return "AuthWxRequest-Authorize"
}

func (agq *AuthGiteeRequest) Login(authCallBack model.AuthCallback) (model.AuthResponse, error) {
	response, err := agq.adq.Login(agq, authCallBack)
	if err != nil {
		return model.AuthResponse{}, err
	}
	return response, nil
}

func (agq *AuthGiteeRequest) Revoke(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (agq *AuthGiteeRequest) Refresh(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (agq *AuthGiteeRequest) getAccessToken(authCallback model.AuthCallback) (model.AuthToken, error) {
	response, err := agq.adq.DoPostAuthorizationCode(authCallback.Code)
	if err != nil {
		return model.AuthToken{}, err
	}
	fmt.Println(response)
	return model.AuthToken{}, nil
}

func (agq *AuthGiteeRequest) getUserInfo(authToken model.AuthToken) model.AuthUser {
	return model.AuthUser{}
}
