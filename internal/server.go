package internal

import (
	"fmt"
	_ "hal9k/internal/command"
	"hal9k/pkg/config"
	"net/http"
)

func NewHttpServer(s *config.ServiceConfig) error {
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", s.Port), nil)
}
