package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gotoeveryone/auth-api/app/config"
	"github.com/gotoeveryone/auth-api/app/domain/entity"
	"github.com/gotoeveryone/auth-api/app/presentation/handler"
)

type stateHandler struct{}

// NewStateHandler is state action handler
func NewStateHandler() handler.StateHandler {
	return &stateHandler{}
}

// Get is get application state
func (h *stateHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, entity.State{
		Status:      "Active",
		Environment: gin.Mode(),
		LogLevel:    config.AppConfig.Log.Level,
		TimeZone:    time.Local.String(),
	})
}

// NoRoute is not found response
func (h *stateHandler) NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, entity.Error{
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
	})
}

// NoMethod is method not allowed response
func (h *stateHandler) NoMethod(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, entity.Error{
		Code:    http.StatusMethodNotAllowed,
		Message: http.StatusText(http.StatusMethodNotAllowed),
	})
}
