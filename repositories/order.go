package repositories

import (
	"waysbucks/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	OrderItem(order models.Order) (models.Order, error)
	DelOrderItem(order models.Order) (models.Order, error)
	GetOrder(ID int) (models.Order, error)
	FindAllOrderUser(ID int) ([]models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)

	GetProductOrder(ID int) (models.Product, error)
	GetTopingOrder(ID []int) ([]models.Toping, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) OrderItem(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error

	return order, err
}
func (r *repository) DelOrderItem(order models.Order) (models.Order, error) {
	err := r.db.Delete(&order).Error

	return order, err
}
func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("Product").Preload("Toping").Preload("Buyyer").First(&order, ID).Error

	return order, err
}
func (r *repository) FindAllOrderUser(ID int) ([]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("Product").Preload("Toping").Preload("Buyyer").Where("buyyer_id = ?", ID).Find(&order).Error

	return order, err
}
func (r *repository) UpdateOrder(order models.Order) (models.Order, error) {
	err := r.db.Save(&order).Error

	return order, err
}
func (r *repository) GetProductOrder(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}
func (r *repository) GetTopingOrder(ID []int) ([]models.Toping, error) {
	var toping []models.Toping
	err := r.db.Find(&toping, ID).Error

	return toping, err
}
