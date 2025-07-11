package api

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Healthcheck struct {
	HealthcheckServices interfaces.IHealthcheckServices
}

func (api *Healthcheck) HealthcheckHandlerHttp(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, msg, nil)
}
