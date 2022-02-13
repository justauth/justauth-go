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

	a := request.AuthDefaultRequest{}

	b := request.AuthGiteeRequest{}

	r1 := a.Authorize("")
	r2 := b.Authorize("")

	response, err := b.Login(model.AuthCallback{
		Code: "85dbb209149edf394c41d6466be8d7e90935fa1761f523444ab00086ef2199ab",
	})
	if err != nil {
		return
	}
	fmt.Println(response.Data)
	fmt.Println(r1)
	fmt.Println(r2)
}
