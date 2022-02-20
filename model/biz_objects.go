package model

type AuthCallback struct {
	Code              string `json:"code"`
	AuthCode          string `json:"auth_code"`
	State             string `json:"cache"`
	AuthorizationCode string `json:"authorization_code"`
	OauthToken        string `json:"oauth_token"`
	OauthVerifier     string `json:"oauth_verifier"`
}

type AuthResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type AuthToken struct {
	AccessToken          string `json:"accessToken"`
	ExpireIn             int64  `json:"expireIn"`
	RefreshToken         string `json:"refreshToken"`
	RefreshTokenExpireIn int    `json:"refreshTokenExpireIn"`
	Uid                  string `json:"uid"`
	OpenId               string `json:"openId"`
	AccessCode           string `json:"accessCode"`
	UnionId              string `json:"unionId"`
	Scope                string `json:"scope"`
	TokenType            string `json:"tokenType"`
}

type AuthConfig struct {
	ClientId               string `json:"clientId"`
	ClientSecret           string `json:"client_secret"`
	RedirectUri            string `json:"redirect_uri"`
	AlipayPublicKey        string `json:"alipay_public_key"`
	UnionId                bool   `json:"union_id"`
	StackOverflowKey       string `json:"stackoverflow_key"`
	AgentId                string `json:"agent_id"`
	DomainPrefix           string `json:"domain_prefix"`
	IgnoreCheckState       bool   `json:"ignore_check_state"`
	Scopes                 string `json:"scopes"`
	DeviceId               string `json:"device_id"`
	ClientOsType           int    `json:"client_os_type"`
	PackId                 string `json:"pack_id"`
	Pkce                   bool   `json:"pkce"`
	AuthServerId           string `json:"auth_server_id"`
	IgnoreCheckRedirectUri bool   `json:"ignore_check_redirect_uri"`
}

type AuthUser struct {
	Uuid        string    `json:"uuid"`
	Username    string    `json:"username"`
	Nickname    string    `json:"nickname"`
	Avatar      string    `json:"avatar"`
	Blog        string    `json:"blog"`
	Company     string    `json:"company"`
	Location    string    `json:"location"`
	Email       string    `json:"email"`
	Remark      string    `json:"remark"`
	Gender      string    `json:"gender"`
	Source      string    `json:"source"`
	AuthToken   AuthToken `json:"clientId"`
	RawUserInfo string    `json:"rawUserInfo"`
}
