package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/edposinitckii-hf/monorepo/svc/echo/internal/service"
)

type echoSvc interface {
	Echo(context.Context, string) (*service.EchoMessage, error)
}

type handler struct {
	svc echoSvc
}

func NewHandler(svc echoSvc) *handler {
	return &handler{svc: svc}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var echoRequest struct {
		Message string `json:"message"`
	}

	if err := fromJSON(r.Body, &echoRequest); err != nil {
		writeError(w, err)
		return
	}
	defer r.Body.Close()

	msg, err := h.svc.Echo(r.Context(), echoRequest.Message)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(&msg)
}

func fromJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func writeError(w http.ResponseWriter, err error) {
	var response struct {
		Error string `json:"error"`
	}
	response.Error = err.Error()

	w.WriteHeader(500)
	_ = json.NewEncoder(w).Encode(&response)
}
