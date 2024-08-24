package transactions

import (
	"gofiber-boilerplate/modules/db"
	"gofiber-boilerplate/modules/jwt"
	"gofiber-boilerplate/modules/transactions/transactionsdto"
	"gofiber-boilerplate/modules/transactions/transactionsmodel"
	"gofiber-boilerplate/modules/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionService interface {
	Init(db db.DbService)
	Topup(userId uuid.UUID, topupDTO *transactionsdto.TopupDTO) (*transactionsdto.TransactionDTO, error)
	Transfer(userId uuid.UUID, transferDTO *transactionsdto.TransferDTO) (*transactionsdto.TransactionDTO, error)
	Payment(userId uuid.UUID, paymentDTO *transactionsdto.PaymentDTO) (*transactionsdto.TransactionDTO, error)
}

type TransactionServiceImpl struct {
	userService user.UserService
	jwtService  jwt.JwtService
	db          *gorm.DB
}

func NewTransactionService(userService user.UserService, jwtService jwt.JwtService) TransactionService {
	return &TransactionServiceImpl{
		userService: userService,
		jwtService:  jwtService,
	}
}

// impl `TransactionService` start

func (service *TransactionServiceImpl) Init(db db.DbService) {
	service.db = db.Default()
}

func (service *TransactionServiceImpl) Topup(userId uuid.UUID, topupDTO *transactionsdto.TopupDTO) (*transactionsdto.TransactionDTO, error) {
	transaction := transactionsmodel.TransactionModel{
		UserID:          userId,
		Status:          transactionsmodel.Success,
		TransactionType: transactionsmodel.TypeCredit,
		Amount:          topupDTO.Amount,
		Action:          transactionsmodel.ActionTopup,
	}

	updatedUser, err := service.userService.AddBalance(userId, topupDTO.Amount)

	if err != nil {
		return nil, err
	}

	transaction.BalanceBefore = updatedUser.Balance - topupDTO.Amount
	transaction.BalanceAfter = updatedUser.Balance
	result := service.db.Create(&transaction)
	dto := transactionsdto.MapTopupResponseDTO(&transaction)
	return dto, result.Error
}

func (service *TransactionServiceImpl) Transfer(userId uuid.UUID, transferDTO *transactionsdto.TransferDTO) (*transactionsdto.TransactionDTO, error) {
	transaction := transactionsmodel.TransactionModel{
		UserID:          userId,
		Status:          transactionsmodel.Pending,
		TransactionType: transactionsmodel.TypeDebit,
		Amount:          transferDTO.Amount,
		Remarks:         transferDTO.Remarks,
		Action:          transactionsmodel.ActionTransfer,
	}

	updatedUser, err := service.userService.AddBalance(userId, -transferDTO.Amount)

	if err != nil {
		return nil, err
	}

	transaction.BalanceBefore = updatedUser.Balance + transferDTO.Amount
	transaction.BalanceAfter = updatedUser.Balance
	result := service.db.Create(&transaction)
	dto := transactionsdto.MapTransferResponseDTO(&transaction)
	return dto, result.Error
}

func (service *TransactionServiceImpl) Payment(userId uuid.UUID, paymentDTO *transactionsdto.PaymentDTO) (*transactionsdto.TransactionDTO, error) {
	transaction := transactionsmodel.TransactionModel{
		UserID:          userId,
		Status:          transactionsmodel.Success,
		TransactionType: transactionsmodel.TypeCredit,
		Amount:          paymentDTO.Amount,
		Remarks:         paymentDTO.Remarks,
		Action:          transactionsmodel.ActionPayment,
	}

	updatedUser, err := service.userService.AddBalance(userId, -paymentDTO.Amount)

	if err != nil {
		return nil, err
	}

	transaction.BalanceBefore = updatedUser.Balance + paymentDTO.Amount
	transaction.BalanceAfter = updatedUser.Balance
	result := service.db.Create(&transaction)
	dto := transactionsdto.MapPaymentResponseDTO(&transaction)
	return dto, result.Error
}
