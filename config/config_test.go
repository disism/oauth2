package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T) {
	if err := InitConfig(); err != nil {
		t.Error(err)
		return
	}
	t.Logf("name: %s", viper.GetString("name"))
}
