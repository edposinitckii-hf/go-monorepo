package transport

import "net/http"

const (
	EchoURL = "/echo"
)

func NewRouter(svc echoSvc) *http.ServeMux {
	m := http.NewServeMux()
	m.Handle(EchoURL, NewHandler(svc))

	return m
}
