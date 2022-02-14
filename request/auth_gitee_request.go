package request

import (
	"errors"
	"fmt"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/model"
	"github.com/tidwall/gjson"
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

	err = checkResponse(response)
	if err != nil {
		return model.AuthToken{}, err
	}

	return model.AuthToken{
		AccessToken:  gjson.Get(response, "access_token").String(),
		RefreshToken: gjson.Get(response, "refresh_token").String(),
		Scope:        gjson.Get(response, "scope").String(),
		TokenType:    gjson.Get(response, "token_type").String(),
		ExpireIn:     gjson.Get(response, "expires_in").Int(),
	}, nil
}

func (giteeRequest *AuthGiteeRequest) getUserInfo(authToken model.AuthToken) (model.AuthUser, error) {
	response, err := giteeRequest.baseAuthRequest.DoGetUserInfo(authToken)
	if err != nil {
		return model.AuthUser{}, err
	}

	err = checkResponse(response)
	if err != nil {
		return model.AuthUser{}, err
	}

	return model.AuthUser{
		Uuid:      gjson.Get(response, "id").String(),
		Username:  gjson.Get(response, "login").String(),
		Avatar:    gjson.Get(response, "avatar_url").String(),
		Blog:      gjson.Get(response, "blog").String(),
		Nickname:  gjson.Get(response, "name").String(),
		Company:   gjson.Get(response, "company").String(),
		Location:  gjson.Get(response, "address").String(),
		Email:     gjson.Get(response, "email").String(),
		Remark:    gjson.Get(response, "bio").String(),
		Gender:    string(config.UserGenderUnknown),
		AuthToken: authToken,
		Source:    string(config.PlatformStrGitee),
	}, nil
}

func checkResponse(response string) error {
	errStr := gjson.Get(response, "error").String()
	if errStr != "" {
		errDesc := gjson.Get(response, "error_description").String()
		return errors.New(errDesc)
	}
	return nil
}
