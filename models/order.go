package models

type Order struct {
	ID       int `json:"id" gorm:"primary_key:auto_increment"`
	Qty      int `json:"qty" gorm:"type:int"`
	SubTotal int `json:"subtotal" gorm:"type:int"`
	//------------------------------------------------------------------//
	TransactionID int                 `json:"transaction_id"`
	Transaction   TransactionResponse `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// -----------------------------------------------------------------//
	ProductID int      `json:"product_id"`
	Product   Product  `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Toping    []Toping `json:"toping" gorm:"many2many:toping_order;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OrderRespon struct {
	ID        int      `json:"id"`
	Qty       int      `json:"qty"`
	SubTotal  int      `json:"subtotal"`
	ProductID int      `json:"product_id"`
	Product   Product  `json:"product"`
	Toping    []Toping `json:"toping_id"`
}

func (OrderRespon) TableName() string {
	return "orders"
}
