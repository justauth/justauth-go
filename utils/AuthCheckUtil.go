package utils

import (
	"github.com/justauth/justauth-go/cache"
	"github.com/justauth/justauth-go/config"
	"github.com/justauth/justauth-go/model"
)

func CheckCode(platform config.PlatformUrl, authCallBack model.AuthCallback) bool {
	if platform.Platform == string(config.PlatformStrTwitter) {
		return true
	}
	code := authCallBack.Code

	if platform.Platform == string(config.PlatformStrAliPay) {
		code = authCallBack.AuthCode
	}
	if platform.Platform == string(config.PlatformStrHuaWei) {
		code = authCallBack.AuthorizationCode
	}
	if code == "" {
		return false
	}
	return true
}

func CheckState(state string, platform config.PlatformUrl, stateCache cache.AuthStateCache) bool {
	if platform.Platform == string(config.PlatformStrTwitter) {
		return true
	}
	if state == "" || !stateCache.ContainsKey(state) {
		return false
	}
	return true
}
