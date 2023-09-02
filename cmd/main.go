package main

import (
	"net/http"
	"rc-tournament/cmd/conf"
	"rc-tournament/domain/consts"
	"rc-tournament/internal/package/handler"
	"rc-tournament/internal/package/service"

	"github.com/labstack/gommon/log"
)

func main() {
	configs, err := conf.GetConfigs()
	if err != nil {
		log.Fatalf("Error loading env values %e", err)
		return
	}

	services := map[string]service.Service{
		consts.Healthcheck: handler.NewAccountsService(configs),
		consts.HelloWorld:  handler.NewHelloWorldService(configs),
	}

	s := service.New(services)

	log.Info("servidor creado correctamente")

	log.Fatal(http.ListenAndServe(configs[consts.PortConst], s.Router()))
}
