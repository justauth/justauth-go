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

	giteeRequest, _ := request.NewAuthGiteeRequest(model.AuthConfig{
		ClientId:     "b7df3559670e1c6a12d04f633984651116c01fdcc7f47aa73c596e1abe55b885",
		ClientSecret: "3c792328a06a44d7aba66dde4d8dff15cbba562aa16b7cd46d6ccaf4d38de2cd",
		RedirectUri:  "http://guaguaguaxia.com/testcallback",
	})
	response, err := giteeRequest.Login(model.AuthCallback{
		Code: "ee95d392d381b2ff84a76f1112cdc35f3584a4a09ed35fc78cfc81c8b090993c",
	})
	if err != nil {
		return
	}
	fmt.Println(response.Data)

}
