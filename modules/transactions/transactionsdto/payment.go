package transactionsdto

import "gofiber-boilerplate/modules/transactions/transactionsmodel"

type PaymentDTO struct {
	Amount  int    `json:"amount" validate:"required,number"`
	Remarks string `json:"remarks" validate:"required"`
}

func MapPaymentResponseDTO(model *transactionsmodel.TransactionModel) *TransactionDTO {
	return &TransactionDTO{
		TopUpID:       &model.ID,
		Amount:        model.Amount,
		Remarks:       &model.Remarks,
		BalanceBefore: model.BalanceBefore,
		BalanceAfter:  model.BalanceAfter,
		CreatedAt:     model.CreatedAt,
	}
}
