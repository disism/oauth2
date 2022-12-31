package oauth2

import (
	"testing"
)

func TestInitDatabase(t *testing.T) {
	if err := InitDatabase(); err != nil {
		t.Error(err)
		return
	}
}
