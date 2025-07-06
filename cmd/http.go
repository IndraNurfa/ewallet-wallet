package cmd

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()

	r.GET("/health", dependency.HealthcheckApi.HealthcheckHandlerHttp)

	walletV1 := r.Group("/wallet/v1")
	walletV1.POST("/", dependency.WalletAPI.Create)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	WalletRepository interfaces.IWalletRepo

	HealthcheckApi interfaces.IHealthcheckAPI

	WalletAPI interfaces.IWalletAPI
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

	return Dependency{
		WalletRepository: walletRepo,
		HealthcheckApi:   healthcheckAPI,
		WalletAPI:        walletAPI,
	}
}
