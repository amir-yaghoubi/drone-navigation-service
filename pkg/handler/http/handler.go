package http

import (
	"drone_navigation_service/pkg/dns"

	"github.com/gin-gonic/gin"
)

func NewHandler(eng *gin.RouterGroup, svc dns.DroneNavigationServicer) *DroneNavigationHandler {
	return &DroneNavigationHandler{eng, svc}
}

type DroneNavigationHandler struct {
	engine *gin.RouterGroup
	svc    dns.DroneNavigationServicer
}

func (h *DroneNavigationHandler) Setup() {
	h.engine.POST("/", h.HandleDNSRequest)
}
