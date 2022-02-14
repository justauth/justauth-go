package request

import (
	"fmt"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/model"
	"io/ioutil"
	"net/http"
)

type BaseAuthRequest struct {
	authConfig  model.AuthConfig
	platformUrl config.PlatformUrl
}

func (baseAuthRequest *BaseAuthRequest) Authorize(state string) string {
	return fmt.Sprintf("%s")
}

func (baseAuthRequest *BaseAuthRequest) Login(request AuthRequest, authCallBack model.AuthCallback) (model.AuthResponse, error) {
	authToken, err := request.getAccessToken(authCallBack)
	if err != nil {
		return model.AuthResponse{}, err
	}
	userInfo, err := request.getUserInfo(authToken)
	if err != nil {
		return model.AuthResponse{}, err
	}
	return model.AuthResponse{
		Code: 200,
		Msg:  "",
		Data: userInfo,
	}, nil
}

func (baseAuthRequest *BaseAuthRequest) Revoke(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (baseAuthRequest *BaseAuthRequest) Refresh(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (baseAuthRequest *BaseAuthRequest) getAccessToken(authCallback model.AuthCallback) (model.AuthToken, error) {
	return model.AuthToken{}, nil
}

func (baseAuthRequest *BaseAuthRequest) getUserInfo(authToken model.AuthToken) model.AuthUser {
	return model.AuthUser{}
}

func (baseAuthRequest *BaseAuthRequest) DoPostAuthorizationCode(code string) (string, error) {
	resp, err := http.Post(baseAuthRequest.getAccessTokenUrl(code), "application/json", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (baseAuthRequest *BaseAuthRequest) DoGetUserInfo(authToken model.AuthToken) (string, error) {
	url := baseAuthRequest.getUserInfoUrl(authToken)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (baseAuthRequest *BaseAuthRequest) getAccessTokenUrl(code string) string {
	return fmt.Sprintf("%s?code=%s&client_id=%s&client_secret=%s&grant_type=authorization_code&redirect_uri=%s",
		baseAuthRequest.platformUrl.AccessTokenUrl, code, baseAuthRequest.authConfig.ClientId, baseAuthRequest.authConfig.ClientSecret, baseAuthRequest.authConfig.RedirectUri)
}

func (baseAuthRequest *BaseAuthRequest) getUserInfoUrl(authToken model.AuthToken) string {
	return fmt.Sprintf("%s?access_token=%s",
		baseAuthRequest.platformUrl.UserInfoUrl, authToken.AccessToken)
}
