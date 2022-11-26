package models

type Cart struct {
	ID    int `json:"id" gorm:"primary_key:auto_increment"`
	Qty   int `json:"qty" gorm:"type:int"`
	Price int `json:"price" gorm:"type:int"`

	// -----------------------------------------------------------------//
	ProductID int     `json:"product_id"`
	Product   Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// =================================================================//
	TopingID []int    `json:"toping_id" form:"toping_id" gorm:"-"`
	Toping   []Toping `json:"toping" gorm:"many2many:toping_cart;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// =================================================================//
	BuyyerID int        `json:"buyyer_id"`
	Buyyer   UserRespon `json:"buyyer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// =================================================================//
}

type CartRespon struct {
	ID        int `json:"id"`
	Qty       int `json:"qty"`
	Price     int `json:"price"`
	ProductID int `json:"product_id"`
	TopingID  int `json:"toping_id"`
	BuyyerID  int `json:"buyyer_id"`
}

func (CartRespon) TableName() string {
	return "carts"
}
