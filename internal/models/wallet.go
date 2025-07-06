package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Wallet struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id" gorm:"uniqueIndex;column:user_id;not null" validate:"required"`
	Balance   float64   `json:"balance" gorm:"column:balance;type:decimal(15,2)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (*Wallet) TableName() string {
	return "wallets"
}

func (l Wallet) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type WalletTransaction struct {
	ID                    int
	WalletID              int       `gorm:"column:wallet_id;not null"`
	Amount                float64   `gorm:"column:amount;type:decimal(15,2);not null"`
	WalletTransactionType string    `gorm:"column:wallet_transaction_type;type:ENUM('CREDIT','DEBIT');not null"`
	Reference             string    `gorm:"uniqueIndex;column:reference;type:varchar(100);not null"`
	CreatedAt             time.Time `json:"-"`
	UpdatedAt             time.Time `json:"-"`
}

func (*WalletTransaction) TableName() string {
	return "wallet_transactions"
}

func (l WalletTransaction) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
