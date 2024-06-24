package server

import (
	"net/http"

	"github.com/armanzhankin/echo-htmx/config"
	"github.com/labstack/echo/v4"
)

func RunServer(cfg config.Config, echo *echo.Echo) error {
	server := &http.Server{
		Addr:    cfg.Host + ":" + cfg.Port,
		Handler: echo,
	}

	return echo.StartServer(server)
}
