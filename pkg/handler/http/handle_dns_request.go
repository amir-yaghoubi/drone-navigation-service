package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary DNS request
// @Description: Calculates datalake location
// @Tags dns
// @Accept json
// @Produce json
// @Param default body DnsRequest true "DNS Request Body"
// @Success 200 {object} DnsResponse
// @Router / [post]
func (h *DroneNavigationHandler) HandleDNSRequest(g *gin.Context) {
	var httpReq DnsRequest
	if err := g.ShouldBindJSON(&httpReq); err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Possibility of errors from this transformation is about zero
	// because we are validation json request, but we check it anyway
	domainReq, err := ToDomainDnsRequest(httpReq)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domainRes, err := h.svc.CalculateLocation(g.Request.Context(), *domainReq)
	if err != nil {
		// We should handle business errors in more practical way
		// and assign appropriate http status code to it
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, ToHttpResponse(*domainRes))
}
