package request

import (
	"errors"
	"fmt"
	"github.com/justauth/justauth-go/cache"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/error"
	"github.com/justauth/justauth-go/model"
	"github.com/tidwall/gjson"
)

type AuthGiteeRequest struct {
	baseAuthRequest BaseAuthRequest
	authConfig      model.AuthConfig
}

type GiteeScope string

const (
	GiteeScopeUserInfo     = "user_info"
	GiteeScopeProjects     = "projects"
	GiteeScopePullRequests = "pull_requests"
	GiteeScopeIssues       = "issues"
	GiteeScopeNotes        = "notes"
	GiteeScopeKeys         = "keys"
	GiteeScopeHook         = "hook"
	GiteeScopeGroups       = "groups"
	GiteeScopeGists        = "gists"
	GiteeScopeEnterprises  = "enterprises"
	GiteeScopeEmails       = "emails"
)

var defaultScope = []string{GiteeScopeUserInfo}

func NewAuthGiteeRequest(authConfig model.AuthConfig, authStateCache ...cache.AuthStateCache) (*AuthGiteeRequest, error) {
	stateCache := authStateCache[0]
	if stateCache == nil {
		stateCache = cache.NewAuthDefaultStateCache()
	}
	err := config.InitConfig()
	if err != nil {
		return nil, err
	}
	isHaveConfig, platformUrl := config.GetRequestUrlListConfig().GetUrlByPlatform(config.PlatformStrGitee)
	if !isHaveConfig {
		return nil, selferror.NotHaveConfigError
	}
	return &AuthGiteeRequest{
		baseAuthRequest: BaseAuthRequest{
			AuthConfig:  authConfig,
			PlatformUrl: platformUrl,
			StateCache:  stateCache,
		},
	}, nil
}

func (giteeRequest *AuthGiteeRequest) Authorize(state string) string {
	authorizeUrl := giteeRequest.baseAuthRequest.getAuthorizeUrl(state)
	scopes := giteeRequest.baseAuthRequest.getScopes(" ", true, defaultScope)
	return fmt.Sprintf("%s?scope=%s", authorizeUrl, scopes)
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
