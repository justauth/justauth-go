package justauth

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/justauth/justauth-go/cache"
	"github.com/justauth/justauth-go/model"
	"github.com/justauth/justauth-go/request"
	"testing"
)

func TestHello(t *testing.T) {
	stateCache := cache.NewAuthDefaultStateCache()
	giteeRequest, _ := request.NewAuthGiteeRequest(model.AuthConfig{
		ClientId:     "b7df3559670e1c6a12d04f633984651116c01fdcc7f47aa73c596e1abe55b885",
		ClientSecret: "3c792328a06a44d7aba66dde4d8dff15cbba562aa16b7cd46d6ccaf4d38de2cd",
		RedirectUri:  "http://guaguaguaxia.com/testcallback",
	}, stateCache)

	state := uuid.New().String()

	giteeRequest.Authorize(state)

	response, err := giteeRequest.Login(model.AuthCallback{
		Code:  "e6d8961a2f4d0e347ed3374503b137e6edd622fc2c4c0e8ef6de7b9b1d0aba08",
		State: state,
	})
	fmt.Println(err)
	fmt.Println(response.Data)

	//datas := url.Values{}
	//datas.Set("111", "222")

	//requestData := utils.HttpRequestData{
	//	Client:   http.DefaultClient,
	//	Method:   http.MethodPost,
	//	ReqUrl:   "http://httpbin.org/post",
	//	Params:   map[string]string{"=": "="},
	//	Data:     map[string]string{"111": "222"},
	//	Headers:  map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
	//	IsEncode: true,
	//	ProxyUrl: "http://60.31.89.41:4231",
	//}
	//resp, err := utils.HttpRequest(requestData)
	//if err != nil {
	//	return
	//}
	//fmt.Println(resp)
}
