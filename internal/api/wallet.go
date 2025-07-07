package api

import (
	"ewallet-wallet/constants"
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WalletAPI struct {
	WalletService interfaces.IWalletService
}

func (api *WalletAPI) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.Wallet
	)

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if req.UserId == 0 {
		log.Error("user_id is empty")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	req.CreatedBy = strconv.Itoa(req.UserId)
	req.UpdatedBy = strconv.Itoa(req.UserId)

	err := api.WalletService.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("failed to create wallet: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.SuccessMessage, req)
}

func (api *WalletAPI) CreditBalance(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.TransactionRequest
	)

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.WalletService.CreditBalance(c.Request.Context(), int(tokenData.UserID), req)
	if err != nil {
		log.Error("failed to credit balance wallet: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.SuccessMessage, resp)
}

func (api *WalletAPI) DebitBalance(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.TransactionRequest
	)

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.WalletService.DebitBalance(c.Request.Context(), int(tokenData.UserID), req)
	if err != nil {
		log.Error("failed to debit balance wallet: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.SuccessMessage, resp)
}

func (api *WalletAPI) GetBalance(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.WalletService.GetBalance(c.Request.Context(), int(tokenData.UserID))
	if err != nil {
		log.Error("failed to get balance wallet: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.SuccessMessage, resp)
}

func (api *WalletAPI) GetWalletHistory(c *gin.Context) {
	var (
		log   = helpers.Logger
		param models.WalletHistoryParam
	)

	if err := c.ShouldBindQuery(&param); err != nil {
		log.Error("failed to parse query params")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if param.WalletTransactionType != "" {
		if param.WalletTransactionType != "CREDIT" && param.WalletTransactionType != "DEBIT" {
			log.Error("invalid walet transaction type")
			helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
			return
		}
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.WalletService.GetWalletHistory(c.Request.Context(), int(tokenData.UserID), param)
	if err != nil {
		log.Error("failed to get balance wallet: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.SuccessMessage, resp)
}
