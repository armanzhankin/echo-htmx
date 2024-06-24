package handler

import (
	"net/http"

	"github.com/armanzhankin/echo-htmx/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Handler struct {
	service   *service.Service
	logger    *zap.SugaredLogger
	jwtSecret string
}

func NewHandler(service *service.Service, logger *zap.SugaredLogger, jwtSecret string) *Handler {
	return &Handler{
		service:   service,
		logger:    logger,
		jwtSecret: jwtSecret,
	}
}

func (h *Handler) Routes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	e.GET("/health", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "everything is ok")
	})

	return e
}
