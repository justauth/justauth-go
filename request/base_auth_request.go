package request

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/justauth/justauth-go/cache"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/error"
	"github.com/justauth/justauth-go/model"
	"github.com/justauth/justauth-go/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type BaseAuthRequest struct {
	AuthConfig  model.AuthConfig
	PlatformUrl config.PlatformUrl
	StateCache  cache.AuthStateCache
}

func (baseAuthRequest *BaseAuthRequest) Login(request AuthRequest, authCallBack model.AuthCallback) (model.AuthResponse, error) {

	checkCodeResult := utils.CheckCode(baseAuthRequest.PlatformUrl, authCallBack)
	if !checkCodeResult {
		return model.AuthResponse{}, selferror.IllegalCodeError
	}

	if !baseAuthRequest.AuthConfig.IgnoreCheckRedirectUri {
		isStateEqual := utils.CheckState(authCallBack.State, baseAuthRequest.PlatformUrl, baseAuthRequest.StateCache)
		if !isStateEqual {
			return model.AuthResponse{}, selferror.IllegalStateError
		}
	}

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

func (baseAuthRequest *BaseAuthRequest) DoGetAuthorizationCode(code string) (string, error) {
	resp, err := http.Get(baseAuthRequest.getAccessTokenUrl(code))
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

func (baseAuthRequest *BaseAuthRequest) DoPostUserInfo(authToken model.AuthToken) (string, error) {
	url := baseAuthRequest.getUserInfoUrl(authToken)
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
		baseAuthRequest.PlatformUrl.AccessTokenUrl, code, baseAuthRequest.AuthConfig.ClientId, baseAuthRequest.AuthConfig.ClientSecret, baseAuthRequest.AuthConfig.RedirectUri)
}

func (baseAuthRequest *BaseAuthRequest) getUserInfoUrl(authToken model.AuthToken) string {
	return fmt.Sprintf("%s?access_token=%s",
		baseAuthRequest.PlatformUrl.UserInfoUrl, authToken.AccessToken)
}

func (baseAuthRequest *BaseAuthRequest) getAuthorizeUrl(state string) string {
	if state == "" {
		state = uuid.New().String()
	}
	baseAuthRequest.StateCache.Cache(state, state)
	return fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&state=%s",
		baseAuthRequest.PlatformUrl.AuthorizeUrl, baseAuthRequest.AuthConfig.ClientId, baseAuthRequest.AuthConfig.RedirectUri, state)
}

func (baseAuthRequest *BaseAuthRequest) getScopes(separator string, isEncode bool, defaultScopes []string) string {
	configScopes := baseAuthRequest.AuthConfig.Scopes
	if len(configScopes) == 0 {
		if len(defaultScopes) == 0 {
			return ""
		}
		configScopes = defaultScopes
	}
	if separator == "" {
		separator = " "
	}
	scopeStr := strings.Join(configScopes, separator)

	if isEncode {
		return url.QueryEscape(scopeStr)
	}
	return scopeStr
}
