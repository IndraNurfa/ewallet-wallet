package cmd

import (
	"ewallet-wallet/external"
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	d := dependencyInject()

	r := gin.Default()

	r.GET("/health", d.HealthcheckApi.HealthcheckHandlerHttp)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", d.WalletAPI.Create)
	walletV1.PUT("/balance/credit", d.ValidateToken, d.WalletAPI.CreditBalance)
	walletV1.PUT("/balance/debit", d.ValidateToken, d.WalletAPI.DebitBalance)
	walletV1.GET("/balance", d.ValidateToken, d.WalletAPI.GetBalance)
	walletV1.GET("/history", d.ValidateToken, d.WalletAPI.GetWalletHistory)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	WalletRepository interfaces.IWalletRepo
	HealthcheckApi   interfaces.IHealthcheckAPI
	WalletAPI        interfaces.IWalletAPI
	External         interfaces.IExternal
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	walletRepo := &repository.WalletRepo{
		DB: helpers.DB,
	}

	walletSvc := &services.WalletService{
		WalletRepo: walletRepo,
	}
	walletAPI := &api.WalletAPI{
		WalletService: walletSvc,
	}

	external := &external.External{}

	return Dependency{
		WalletRepository: walletRepo,
		HealthcheckApi:   healthcheckAPI,
		WalletAPI:        walletAPI,
		External:         external,
	}
}
