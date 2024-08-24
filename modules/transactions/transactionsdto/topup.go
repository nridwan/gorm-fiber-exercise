package transactionsdto

import "gofiber-boilerplate/modules/transactions/transactionsmodel"

type TopupDTO struct {
	Amount int `json:"amount" validate:"required,number"`
}

func MapTopupResponseDTO(model *transactionsmodel.TransactionModel) *TransactionDTO {
	return &TransactionDTO{
		TopUpID:       &model.ID,
		Amount:        model.Amount,
		BalanceBefore: model.BalanceBefore,
		BalanceAfter:  model.BalanceAfter,
		CreatedAt:     model.CreatedAt,
	}
}
