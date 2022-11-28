package models

type Transaction struct {
	ID        int        `json:"id" gorm:"primary_key:auto_increment"`
	Name      string     `json:"name" gorm:"type:text"`
	Address   string     `json:"address" gorm:"type:text"`
	Total     int        `json:"total" gorm:"type:int"`
	Status    string     `json:"status" gorm:"type:text"`
	OrderID   int        `json:"order_id"`
	Order     []Order    `json:"order" gorm:"many2many:transaction_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccountID int        `json:"account_id"`
	Account   UserRespon `json:"account" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
