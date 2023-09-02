package handler

import (
	"net/http"
	"rc-tournament/domain/entities"
	"rc-tournament/internal/package/service"

	"github.com/labstack/gommon/log"
)

type helloWorld struct {
	configs map[string]string
}

func NewHelloWorldService(configs map[string]string) service.Service {
	return &helloWorld{configs: configs}
}

func (h *helloWorld) Resolver(w http.ResponseWriter, r *http.Request) {
	log.Info("hello world solicited")
	entities.AccountsAnswer(http.StatusAccepted, "hello world", w)
}
