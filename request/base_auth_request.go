package request

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/justauth/justauth-go/cache"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/error"
	"github.com/justauth/justauth-go/model"
	"github.com/justauth/justauth-go/utils"
	"net/http"
	"net/url"
	"strings"
)

type BaseAuthRequest struct {
	AuthConfig  model.AuthConfig
	PlatformUrl config.PlatformUrl
	StateCache  cache.AuthStateCache
	HttpClient  *http.Client
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

func (baseAuthRequest *BaseAuthRequest) AcquireAuthorizationCode(httpMethod, code string) (string, error) {
	response, err := utils.HttpRequest(utils.HttpRequestData{
		Method: httpMethod,
		ReqUrl: baseAuthRequest.getAccessTokenUrl(code),
		Client: baseAuthRequest.HttpClient,
	})
	if err != nil {
		return "", err
	}
	return response, nil
}

func (baseAuthRequest *BaseAuthRequest) AcquireUserInfo(httpMethod string, authToken model.AuthToken) (string, error) {
	response, err := utils.HttpRequest(utils.HttpRequestData{
		Method: httpMethod,
		ReqUrl: baseAuthRequest.getUserInfoUrl(authToken),
		Client: baseAuthRequest.HttpClient,
	})
	if err != nil {
		return "", err
	}
	return response, nil
}

func (baseAuthRequest *BaseAuthRequest) AcquireRevoke(httpMethod string, authToken model.AuthToken) (string, error) {
	response, err := utils.HttpRequest(utils.HttpRequestData{
		Method: httpMethod,
		ReqUrl: baseAuthRequest.getRevokeUrl(authToken),
		Client: baseAuthRequest.HttpClient,
	})
	if err != nil {
		return "", err
	}
	return response, nil
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

func (baseAuthRequest *BaseAuthRequest) getRevokeUrl(authToken model.AuthToken) string {
	return fmt.Sprintf("%s?access_token=%s",
		baseAuthRequest.PlatformUrl.RevokeUrl, authToken.AccessToken)
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
