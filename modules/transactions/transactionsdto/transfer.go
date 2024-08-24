package transactionsdto

import (
	"gofiber-boilerplate/modules/transactions/transactionsmodel"

	"github.com/google/uuid"
)

type TransferDTO struct {
	TargetUser  uuid.UUID                           `json:"target_user" validate:"required,uuid4"`
	Amount      int                                 `json:"amount" validate:"required,number"`
	Remarks     string                              `json:"remarks" validate:"required"`
	Transaction *transactionsmodel.TransactionModel `json:"-"`
}

func MapTransferResponseDTO(model *transactionsmodel.TransactionModel) *TransactionDTO {
	return &TransactionDTO{
		TopUpID:       &model.ID,
		Amount:        model.Amount,
		Remarks:       &model.Remarks,
		BalanceBefore: model.BalanceBefore,
		BalanceAfter:  model.BalanceAfter,
		CreatedAt:     model.CreatedAt,
	}
}
