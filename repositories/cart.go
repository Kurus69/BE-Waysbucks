package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddCart(cart models.Cart) (models.Cart, error)
	DelCart(cart models.Cart) (models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	GetProductCart(ID int) (models.Product, error)
	GetTopingCart(ID []int) ([]models.Toping, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}
func (r *repository) DelCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}
func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").Preload("Toping").Preload("Buyyer").First(&cart, ID).Error

	return cart, err
}
func (r *repository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}

func (r *repository) GetProductCart(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Preload("Seller").First(&product, ID).Error

	return product, err
}

func (r *repository) GetTopingCart(ID []int) ([]models.Toping, error) {
	var toping []models.Toping
	err := r.db.Preload("Seller").Find(&toping, ID).Error

	return toping, err
}
