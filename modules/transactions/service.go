package transactions

import (
	"context"
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
	Destroy()
	Topup(userId uuid.UUID, topupDTO *transactionsdto.TopupDTO) (*transactionsdto.TransactionDTO, error)
	Transfer(userId uuid.UUID, transferDTO *transactionsdto.TransferDTO) (*transactionsdto.TransactionDTO, error)
	Payment(userId uuid.UUID, paymentDTO *transactionsdto.PaymentDTO) (*transactionsdto.TransactionDTO, error)
	Report(userId uuid.UUID) ([]*transactionsdto.TransactionDTO, error)
}

type transactionServiceImpl struct {
	userService user.UserService
	jwtService  jwt.JwtService
	db          *gorm.DB
	queue       *chan *transactionsdto.TransferDTO
	runner      *context.CancelFunc
}

func NewTransactionService(userService user.UserService, jwtService jwt.JwtService) TransactionService {
	queue := make(chan *transactionsdto.TransferDTO, 10)
	return &transactionServiceImpl{
		userService: userService,
		jwtService:  jwtService,
		queue:       &queue,
	}
}

// impl `TransactionService` start

func (service *transactionServiceImpl) Init(db db.DbService) {
	service.db = db.Default()
	runnerCtx, cancel := context.WithCancel(context.Background())
	runner := func() {
		defer close(*service.queue)
		for {
			select {
			case <-runnerCtx.Done():
				return
			case dto := <-*service.queue:
				service.receiveTransfer(dto)
			}
		}
	}
	go runner()

	service.runner = &cancel
}

func (service *transactionServiceImpl) Destroy() {
	(*service.runner)()
}

func (service *transactionServiceImpl) Topup(userId uuid.UUID, topupDTO *transactionsdto.TopupDTO) (*transactionsdto.TransactionDTO, error) {
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

func (service *transactionServiceImpl) Transfer(userId uuid.UUID, transferDTO *transactionsdto.TransferDTO) (*transactionsdto.TransactionDTO, error) {
	transaction := transactionsmodel.TransactionModel{
		UserID:          userId,
		Status:          transactionsmodel.Pending,
		TransactionType: transactionsmodel.TypeDebit,
		Amount:          transferDTO.Amount,
		Remarks:         transferDTO.Remarks,
		Action:          transactionsmodel.ActionTransfer,
	}
	transferDTO.Transaction = &transaction

	updatedUser, err := service.userService.AddBalance(userId, -transferDTO.Amount)

	if err != nil {
		return nil, err
	}

	transaction.BalanceBefore = updatedUser.Balance + transferDTO.Amount
	transaction.BalanceAfter = updatedUser.Balance
	result := service.db.Create(&transaction)
	(*service.queue) <- transferDTO
	dto := transactionsdto.MapTransferResponseDTO(&transaction)
	return dto, result.Error
}

func (service *transactionServiceImpl) Payment(userId uuid.UUID, paymentDTO *transactionsdto.PaymentDTO) (*transactionsdto.TransactionDTO, error) {
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

func (service *transactionServiceImpl) Report(userId uuid.UUID) ([]*transactionsdto.TransactionDTO, error) {
	transactions := []transactionsmodel.TransactionModel{}
	query := service.db.Model(transactions)

	// Perform count and find concurrently using goroutines
	err := query.Where("user_id = ?", userId).Find(&transactions).Error

	if err != nil {
		return nil, err
	}

	response := make([]*transactionsdto.TransactionDTO, len(transactions))
	for i, v := range transactions {
		response[i] = transactionsdto.MapTransactionModelToDTO(&v)
	}

	return response, nil
}

// impl `TransactionService` end

func (service *transactionServiceImpl) receiveTransfer(transferDTO *transactionsdto.TransferDTO) error {
	_, err := service.userService.AddBalance(transferDTO.TargetUser, transferDTO.Amount)

	if err != nil {
		service.userService.AddBalance(transferDTO.Transaction.UserID, transferDTO.Amount)
		transferDTO.Transaction.Status = transactionsmodel.Failed
	} else {
		transferDTO.Transaction.Status = transactionsmodel.Success
	}
	return service.db.Save(transferDTO.Transaction).Error
}
