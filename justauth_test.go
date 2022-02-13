package justauth

import (
	"fmt"
	"github.com/justauth/justauth-go/model"
	"github.com/justauth/justauth-go/request"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestHello(t *testing.T) {
	want := "Hello, justauth."

	assert.Equal(t, want, Hello())

	//b := request.AuthGiteeRequest{}
	//
	//r2 := b.Authorize("")
	//
	//response, err := b.Login(model.AuthCallback{
	//	Code: "85dbb209149edf394c41d6466be8d7e90935fa1761f523444ab00086ef2199ab",
	//})
	//if err != nil {
	//	return
	//}
	//fmt.Println(response.Data)
	//fmt.Println(r2)

	giteeRequest, _ := request.NewAuthGiteeRequest(model.AuthConfig{
		ClientId:     "b7df3559670e1c6a12d04f633984651116c01fdcc7f47aa73c596e1abe55b885",
		ClientSecret: "3c792328a06a44d7aba66dde4d8dff15cbba562aa16b7cd46d6ccaf4d38de2cd",
		RedirectUri:  "http://guaguaguaxia.com/testcallback",
	})
	response, err := giteeRequest.Login(model.AuthCallback{
		Code: "e62a099bd49992f0b964be8f1cf685f8955d132cbb65b7531d5933a43fb12541",
	})
	if err != nil {
		return
	}
	fmt.Println(response.Data)

}
