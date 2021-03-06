package interfaces

import (
	"github.com/aswinda/loket-backend-test/models"
)

type ITransactionService interface {
	GetTransactionDetail(transactionId int) (models.TransactionInfoViewModel, error)
	CreateTransaction(body models.TransactionModel) (models.TransactionModel, error)
	PurchaseTransaction(body models.TransactionPurchase) (models.TransactionModel, error)
}
