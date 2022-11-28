package models

type Order struct {
	ID       int `json:"id" gorm:"primary_key:auto_increment"`
	Qty      int `json:"qty" gorm:"type:int"`
	SubTotal int `json:"subtotal" gorm:"type:int"`

	// -----------------------------------------------------------------//
	ProductID int     `json:"product_id"`
	Product   Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// =================================================================//
	// TopingID []int    `json:"toping_id" form:"toping_id" gorm:"-"`
	Toping []Toping `json:"toping" gorm:"many2many:toping_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// =================================================================//
	BuyyerID int        `json:"buyyer_id"`
	Buyyer   UserRespon `json:"buyyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// =================================================================//
}

type OrderRespon struct {
	ID        int   `json:"id"`
	Qty       int   `json:"qty"`
	SubTotal  int   `json:"subtotal"`
	ProductID int   `json:"product_id"`
	TopingID  []int `json:"toping_id"`
	BuyyerID  int   `json:"buyyer_id"`
}

func (OrderRespon) TableName() string {
	return "orders"
}
