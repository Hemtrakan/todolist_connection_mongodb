package healthCheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HealthCheckHandler interface {
	GetAlive(c *gin.Context)
}

type handler struct {
}

func NewHandler() HealthCheckHandler {
	return &handler{}
}

// @Description  check alive api
// @Tags         healthCheck
// @Accept       json
// @Produce      json
// @Success      200  {object}  GetAliveResponse
// @Failure      400  {object}  httpResponse.HttpResponse
// @Failure      404  {object}  httpResponse.HttpResponse
// @Failure      500  {object}  httpResponse.HttpResponse
// @Router       /api/health-check [get]
func (h *handler) GetAlive(c *gin.Context) {
	response := GetAliveResponse{
		Api:         viper.GetString("app.name"),
		Version:     viper.GetString("app.version"),
		Environment: viper.GetString("app.environment"),
	}
	c.JSON(http.StatusOK, response)
}
