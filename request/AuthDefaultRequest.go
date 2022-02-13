package request

import (
	"fmt"
	"github.com/justauth/justauth-go/model"
	"io/ioutil"
	"net/http"
)

type AuthDefaultRequest struct {
	authConfig model.AuthConfig
}

func (adq *AuthDefaultRequest) Authorize(state string) string {
	return fmt.Sprintf("%s")
}

func (adq *AuthDefaultRequest) Login(request AuthRequest, authCallBack model.AuthCallback) (model.AuthResponse, error) {
	authToken, err := request.getAccessToken(authCallBack)
	if err != nil {
		return model.AuthResponse{}, err
	}
	userInfo := request.getUserInfo(authToken)
	return model.AuthResponse{
		Code: 200,
		Msg:  "",
		Data: userInfo,
	}, nil
}

func (adq *AuthDefaultRequest) Revoke(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (adq *AuthDefaultRequest) Refresh(authToken model.AuthToken) model.AuthResponse {
	return model.AuthResponse{}
}

func (adq *AuthDefaultRequest) getAccessToken(authCallback model.AuthCallback) (model.AuthToken, error) {
	return model.AuthToken{}, nil
}

func (adq *AuthDefaultRequest) getUserInfo(authToken model.AuthToken) model.AuthUser {
	return model.AuthUser{}
}

func (adq *AuthDefaultRequest) DoPostAuthorizationCode(code string) (string, error) {
	clientId := "b7df3559670e1c6a12d04f633984651116c01fdcc7f47aa73c596e1abe55b885"
	clientSecret := "3c792328a06a44d7aba66dde4d8dff15cbba562aa16b7cd46d6ccaf4d38de2cd"
	url := fmt.Sprintf("https://gitee.com/oauth/token?grant_type=authorization_code"+
		"&code=%s&client_id=%s&redirect_uri=%s&client_secret=%s", code, clientId, "http://guaguaguaxia.com/testcallback", clientSecret)
	fmt.Println(url)
	resp, err := http.Post(url, "application/json", nil)
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
