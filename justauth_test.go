package justauth

import (
	"github.com/justauth/justauth-go/request"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestHello(t *testing.T) {
	want := "Hello, justauth."

	assert.Equal(t, want, Hello())

	authWxRequest := &request.AuthWxRequest{}

	authWxRequest.Play(authWxRequest)
}
