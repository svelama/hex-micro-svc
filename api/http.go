package api

import (
	"net/http"

	"github.com/svelama/microsvc/shortener"
)


type RedirectHandler interface{
	Get(w http.ResponseWriter, req *http.Request)
	Post(w http.ResponseWriter, req *http.Request)
}

type handler struct{
	redirectService shortener.RedirectService
}

func NewHandler(redirectService shortener.RedirectService) RedirectHandler{
	return &handler{ redirectService }
}

func (h handler) Get(w http.ResponseWriter, req *http.Request) {
	panic("not implemented") // TODO: Implement
}

func (h handler) Post(w http.ResponseWriter, req *http.Request) {
	panic("not implemented") // TODO: Implement
}

