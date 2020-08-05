package justauth

import "testing"
import "github.com/stretchr/testify/assert"

func TestHello(t *testing.T) {
	want := "Hello, justauth."

	assert.Equal(t, want, Hello())
}
