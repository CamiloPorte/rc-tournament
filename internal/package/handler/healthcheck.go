package handler

import (
	"net/http"
	"rc-tournament/domain/entities"
	"rc-tournament/internal/package/service"
)

type accounts struct {
	configs map[string]string
}

func NewAccountsService(configs map[string]string) service.Service {
	return &accounts{configs: configs}
}

func (a *accounts) Resolver(w http.ResponseWriter, r *http.Request) {
	entities.AccountsAnswer(http.StatusAccepted, "Connection healthy", w)
}
