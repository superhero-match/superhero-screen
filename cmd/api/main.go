package main

import (
	"github.com/superheroville-municipality/cmd/api/controller"
	"github.com/superheroville-municipality/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	ctrl, err := controller.NewController(cfg)
	if err != nil {
		panic(err)
	}

	r := ctrl.RegisterRoutes()

	err = r.RunTLS(
			cfg.App.Port,
			cfg.App.CertFile,
			cfg.App.KeyFile,
		)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = ctrl.Producer.Close()
	}()
}
