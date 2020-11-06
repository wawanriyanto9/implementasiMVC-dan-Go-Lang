package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imadedwisuryadinata/implementasiMvcGolang/app/model"
	"github.com/imadedwisuryadinata/implementasiMvcGolang/app/utils"
	"gorm.io/gorm"
)

type TransactionController struct {
	DB *gorm.DB
}

func (ctrl TransactionController) Transfer(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Transfer(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknown error", http.StatusInternalServerError)
		return
	}
	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

func (ctrl TransactionController) Withdraw(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Withdraw(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknown error", http.StatusInternalServerError)
		return
	}
	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}

func (ctrl TransactionController) Deposit(ctx *gin.Context) {
	transactionModel := model.TransactionModel{
		DB: ctrl.DB,
	}
	var trx model.Transaction

	err := ctx.Bind(&trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	flag, err := transactionModel.Deposit(trx)
	if err != nil {
		utils.WrapAPIError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	if !flag {
		utils.WrapAPIError(ctx, "Unknown error", http.StatusInternalServerError)
		return
	}
	utils.WrapAPISuccess(ctx, "success", http.StatusOK)
	return
}
