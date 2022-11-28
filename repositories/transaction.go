package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	AddTransaction(transaction models.Transaction) (models.Transaction, error)
	GetOrderByUser(ID int) ([]models.Order, error)
	CancelTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
}

func RepoTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}
func (r *repository) GetOrderByUser(ID int) ([]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("Product").Preload("Toping").Preload("Buyyer").Where("buyyer_id = ?", ID).Find(&order).Error

	return order, err
}
func (r *repository) CancelTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Order").Preload("Account").First(&transaction, ID).Error

	return transaction, err
}
func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}
