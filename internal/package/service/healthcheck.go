package service

import (
	"net/http"
	"rc-tournament/domain/consts"

	"github.com/gorilla/mux"
)

type Service interface {
	Resolver(w http.ResponseWriter, r *http.Request)
}

type api struct {
	router  http.Handler
	service map[string]Service
}

type Server interface {
	Router() http.Handler
}

func New(service map[string]Service) Server {
	a := &api{
		service: service,
	}

	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", a.service[consts.Healthcheck].Resolver).Methods(http.MethodGet)
	r.HandleFunc("/", a.service[consts.HelloWorld].Resolver).Methods(http.MethodGet)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
