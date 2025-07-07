package models

import (
	"time"
)

type Wallet struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id" gorm:"uniqueIndex;column:user_id;not null" validate:"required"`
	Balance   float64   `json:"balance" gorm:"column:balance;type:decimal(15,2)"`
	CreatedAt time.Time `json:"-"`
	CreatedBy string    `gorm:"column:created_by;not null"`
	UpdatedAt time.Time `json:"-"`
	UpdatedBy string    `gorm:"column:updated_by;not null"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

type WalletTransaction struct {
	ID                    int       `json:"-"`
	WalletID              int       `json:"wallet_id" gorm:"column:wallet_id;not null"`
	Amount                float64   `json:"amount" gorm:"column:amount;type:decimal(15,2);not null"`
	WalletTransactionType string    `json:"wallet_transaction_type" gorm:"column:wallet_transaction_type;type:ENUM('CREDIT','DEBIT');not null"`
	Reference             string    `json:"reference" gorm:"uniqueIndex;column:reference;type:varchar(100);not null"`
	CreatedAt             time.Time `json:"date"`
	CreatedBy             string    `json:"-" gorm:"column:created_by;not null"`
	UpdatedAt             time.Time `json:"-"`
	UpdatedBy             string    `json:"-" gorm:"column:updated_by;not null"`
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}

type WalletHistoryParam struct {
	Page                  int    `form:"page"`
	Limit                 int    `form:"limit"`
	WalletTransactionType string `form:"wallet_transaction_type"`
}
