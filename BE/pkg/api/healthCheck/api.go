package healthCheck

import (
	"github.com/gin-gonic/gin"
)

func APIHealthCheckServices(r *gin.RouterGroup) {
	healthCheckHandler := NewHandler()
	r.GET("/health-check", healthCheckHandler.GetAlive)

}
