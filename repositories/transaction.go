package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository interface {
	AddTransaction(transaction models.Transaction) (models.Transaction, error)
	CancelTransaction(transaction models.Transaction) (models.Transaction, error)

	//=================================================================================
	GetOrderTrans(ID int) ([]models.Order, error)
	GetTransaction(ID int) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	AllTransaction() ([]models.Transaction, error)
	AllTransactionUser(ID int) ([]models.Transaction, error)
	UpdateTransactionUser(status string, ID int) error
}

func RepoTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}
func (r *repository) CancelTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Order").Preload("User").First(&transaction, ID).Error

	return transaction, err
}
func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}
func (r *repository) AllTransaction() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Order").Preload("User").Where("status != ? ", "Pending").Find(&transaction).Error

	return transaction, err
}
func (r *repository) GetOrderTrans(ID int) ([]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("Product").Preload("Toping").Where("transaction_id = ?", ID).Find(&order).Error

	return order, err
}
func (r *repository) AllTransactionUser(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Order.Product").Preload("Order.Toping").Preload("User").Preload(clause.Associations).Where("status != ? AND user_id = ?", "Pending", ID).Find(&transaction).Error

	return transaction, err
}
func (r *repository) UpdateTransactionUser(status string, ID int) error {
	var transaction models.Transaction
	r.db.Preload("Order.Product").First(&transaction, ID)
	transaction.Status = status
	err := r.db.Debug().Save(&transaction).Error
	return err
}
