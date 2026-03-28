package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"rest-api-2/internal/usecases"
)

type Handlers struct {
	usecases *usecases.UseCases
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{usecases: useCases}
}

func (h Handlers) Listen(port int) error {
	http.HandleFunc("/users", h.registerUserEndpoints)
	slog.Info("Server correndo na porta: " + fmt.Sprintf("%v", port))
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
