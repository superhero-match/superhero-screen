package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/superhero-screen/internal/config"
	"github.com/superhero-screen/internal/es"
)

const (
	timeFormat = "2006-01-02T15:04:05"
)

// Controller holds the controller data.
type Controller struct {
	ES       *es.ES
}

// NewController returns new controller.
func NewController(cfg *config.Config) (ctrl *Controller, err error) {
	e, err := es.NewES(cfg)
	if err != nil {
		return nil, err
	}

	return &Controller{
		ES:       e,
	}, nil
}

// RegisterRoutes registers all the superhero_screen API routes.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/superhero_screen")

	// sr.Use(c.Authorize)

	sr.POST("/check_email", ctl.CheckEmail)

	return router
}
