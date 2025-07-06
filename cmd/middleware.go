package cmd

import (
	"ewallet-wallet/external"
	"ewallet-wallet/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) ValidateToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized empty", nil)
		c.Abort()
		return
	}

	tokenData, err := external.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Error(err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized empty", nil)
		c.Abort()
		return
	}

	c.Set("token", tokenData)

	c.Next()
}
