package transactionsdto

import (
	"gofiber-boilerplate/modules/transactions/transactionsmodel"
	"time"

	"github.com/google/uuid"
)

type TransactionDTO struct {
	TransferID      *uuid.UUID                  `json:"transfer_id,omitempty"`
	PaymentID       *uuid.UUID                  `json:"payment_id,omitempty"`
	TopUpID         *uuid.UUID                  `json:"top_up_id,omitempty"`
	Status          transactionsmodel.TrxStatus `json:"status"`
	UserID          uuid.UUID                   `json:"user_id"`
	TransactionType transactionsmodel.TrxType   `json:"transaction_type"`
	Amount          int                         `json:"amount"`
	Remarks         *string                     `json:"remarks,omitempty"`
	BalanceBefore   int                         `json:"balance_before"`
	BalanceAfter    int                         `json:"balance_after"`
	CreatedAt       *time.Time                  `json:"created_date,omitempty"`
}

func MapTransactionModelToDTO(model *transactionsmodel.TransactionModel) *TransactionDTO {
	var transferID *uuid.UUID
	var paymentID *uuid.UUID
	var topUpID *uuid.UUID

	switch model.Action {
	case transactionsmodel.ActionPayment:
		paymentID = &model.ID
	case transactionsmodel.ActionTransfer:
		transferID = &model.ID
	case transactionsmodel.ActionTopup:
		topUpID = &model.ID
	}

	return &TransactionDTO{
		TransferID:      transferID,
		PaymentID:       paymentID,
		TopUpID:         topUpID,
		Status:          model.Status,
		UserID:          model.UserID,
		TransactionType: model.TransactionType,
		Amount:          model.Amount,
		Remarks:         &model.Remarks,
		BalanceBefore:   model.BalanceBefore,
		BalanceAfter:    model.BalanceAfter,
		CreatedAt:       model.CreatedAt,
	}
}
