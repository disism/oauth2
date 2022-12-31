package oauth2

import "github.com/disism/oauth2/config"

func Run() error {
	if err := config.InitConfig(); err != nil {
		return err
	}

	if err := InitDatabase(); err != nil {
		return err
	}

	if err := Router().Run(":9096"); err != nil {
		return err
	}
	return nil
}
