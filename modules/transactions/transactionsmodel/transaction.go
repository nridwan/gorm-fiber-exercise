package transactionsmodel

import (
	"gofiber-boilerplate/modules/user/usermodel"
	"time"

	"github.com/google/uuid"
)

type TrxStatus string
type TrxAction string
type TrxType string

const (
	Success TrxStatus = "SUCCESS"
	Pending TrxStatus = "PENDING"
	Failed  TrxStatus = "FAILED"

	TypeDebit  TrxType = "DEBIT"
	TypeCredit TrxType = "CREDIT"

	ActionTopup    TrxAction = "TOPUP"
	ActionTransfer TrxAction = "TRANSFER"
	ActionPayment  TrxAction = "PAYMENT"
)

type TransactionModel struct {
	ID              uuid.UUID           `json:"id" gorm:"type:uuid;not null;primaryKey;default:uuid_generate_v4()"`
	Status          TrxStatus           `json:"status" gorm:"not null;"`
	UserID          uuid.UUID           `json:"user_id" gorm:"not null;"`
	User            usermodel.UserModel `json:"user" gorm:"not null;"`
	TransactionType TrxType             `json:"transaction_type" gorm:"not null;"`
	Amount          int                 `json:"amount" gorm:"not null;"`
	Remarks         string              `json:"remarks" gorm:"not null;"`
	BalanceBefore   int                 `json:"balance_before" gorm:"not null;"`
	BalanceAfter    int                 `json:"balance_after" gorm:"not null;"`
	Action          TrxAction           `json:"-"`
	CreatedAt       *time.Time          `json:"created_date,omitempty" gorm:"not null;"`
}

func (TransactionModel) TableName() string {
	return "transactions"
}
