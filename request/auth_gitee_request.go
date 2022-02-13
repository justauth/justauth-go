package request

import (
	"errors"
	"fmt"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/model"
)

type AuthGiteeRequest struct {
	baseAuthRequest BaseAuthRequest
	authConfig      model.AuthConfig
}

func NewAuthGiteeRequest(authConfig model.AuthConfig) (*AuthGiteeRequest, error) {
	err := config.InitConfig()
	if err != nil {
		return nil, err
	}
	isHaveConfig, platformUrl := config.GetRequestUrlListConfig().GetUrlByPlatform(config.PlatformStrGitee)
	if !isHaveConfig {
		return nil, errors.New(fmt.Sprintf("%s not have config", "NewAuthGiteeRequest"))
	}
	return &AuthGiteeRequest{
		baseAuthRequest: BaseAuthRequest{
			authConfig:  authConfig,
			platformUrl: platformUrl,
		},
	}, nil
}

func (giteeRequest *AuthGiteeRequest) Authorize(state string) string {
	return "AuthWxRequest-Authorize"
}

func (giteeRequest *AuthGiteeRequest) Login(authCallBack model.AuthCallback) (model.AuthResponse, error) {
	response, err := giteeRequest.baseAuthRequest.Login(giteeRequest, authCallBack)
	if err != nil {
		return model.AuthResponse{}, err
	}
	return response, nil
}

func (giteeRequest *AuthGiteeRequest) Revoke(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (giteeRequest *AuthGiteeRequest) Refresh(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (giteeRequest *AuthGiteeRequest) getAccessToken(authCallback model.AuthCallback) (model.AuthToken, error) {
	response, err := giteeRequest.baseAuthRequest.DoPostAuthorizationCode(authCallback.Code)
	if err != nil {
		return model.AuthToken{}, err
	}
	fmt.Println(response)
	return model.AuthToken{}, nil
}

func (giteeRequest *AuthGiteeRequest) getUserInfo(authToken model.AuthToken) model.AuthUser {
	return model.AuthUser{}
}
