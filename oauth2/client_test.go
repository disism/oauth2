package oauth2

import (
	"testing"
)

func TestRandToken(t *testing.T) {
	x := RandSecrets(24)
	t.Log(x)
}
