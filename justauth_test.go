package justauth

import (
	"fmt"
	"github.com/justauth/justauth-go/request"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestHello(t *testing.T) {
	want := "Hello, justauth."

	assert.Equal(t, want, Hello())

	a := request.AuthDefaultRequest{}

	b := request.AuthWxRequest{}

	r1 := a.Authorize("")
	r2 := b.Authorize("")

	fmt.Println(r1)
	fmt.Println(r2)
}
