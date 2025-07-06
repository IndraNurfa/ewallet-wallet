package interfaces

import (
	"context"
	"ewallet-wallet/internal/models"

	"github.com/gin-gonic/gin"
)

type IWalletAPI interface {
	Create(c *gin.Context)
}

type IWalletService interface {
	Create(ctx context.Context, wallet *models.Wallet) error
}

type IWalletRepo interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
}
