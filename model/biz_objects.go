package model

type AuthCallback struct {
	Code              string `json:"code"`
	AuthCode          string `json:"auth_code"`
	State             string `json:"state"`
	AuthorizationCode string `json:"authorization_code"`
	OauthToken        string `json:"oauth_token"`
	OauthVerifier     string `json:"oauth_verifier"`
}

type AuthResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type AuthToken struct {
	AccessToken          string `json:"accessToken"`
	ExpireIn             int    `json:"expireIn"`
	RefreshToken         string `json:"refreshToken"`
	RefreshTokenExpireIn int    `json:"refreshTokenExpireIn"`
	Uid                  string `json:"uid"`
	OpenId               string `json:"openId"`
	AccessCode           string `json:"accessCode"`
	UnionId              string `json:"unionId"`
}
