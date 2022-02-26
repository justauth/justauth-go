package justauth

import (
	"fmt"
	"github.com/justauth/justauth-go/utils"
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	//want := "Hello, justauth."
	//
	//assert.Equal(t, want, Hello())
	//
	//giteeRequest, _ := request.NewAuthGiteeRequest(model.AuthConfig{
	//	ClientId:     "b7df3559670e1c6a12d04f633984651116c01fdcc7f47aa73c596e1abe55b885",
	//	ClientSecret: "3c792328a06a44d7aba66dde4d8dff15cbba562aa16b7cd46d6ccaf4d38de2cd",
	//	RedirectUri:  "http://guaguaguaxia.com/testcallback",
	//})
	//response, err := giteeRequest.Login(model.AuthCallback{
	//	Code: "ee95d392d381b2ff84a76f1112cdc35f3584a4a09ed35fc78cfc81c8b090993c",
	//})
	//if err != nil {
	//	return
	//}
	//fmt.Println(response.Data)

	//datas := url.Values{}
	//datas.Set("111", "222")

	requestData := utils.HttpRequestData{
		Client:   http.DefaultClient,
		Method:   http.MethodPost,
		ReqUrl:   "http://httpbin.org/post",
		Params:   map[string]string{"=": "="},
		Data:     map[string]string{"111": "222"},
		Headers:  map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
		IsEncode: true,
		ProxyUrl: "http://60.31.89.41:4231",
	}
	resp, err := utils.HttpRequest(requestData)
	if err != nil {
		return
	}
	fmt.Println(resp)
}
