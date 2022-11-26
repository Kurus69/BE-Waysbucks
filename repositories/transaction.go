package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Checkout(transaction *models.Transaction) (*models.Transaction, error)
	GetCartByUser(ID int) ([]models.Cart, error)
}

func RepoTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Checkout(transaction *models.Transaction) (*models.Transaction, error) {
	err := r.db.Create(transaction).Error

	return transaction, err
}

func (r *repository) GetCartByUser(ID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Buyyer").Where("buyyer_id = ?", ID).Find(&cart).Error

	return cart, err
}
