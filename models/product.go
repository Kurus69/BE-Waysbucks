package models

type Product struct {
	ID       int        `json:"id" gorm:"primary_key:auto_increment"`
	Title    string     `json:"title" gorm:"type:text"`
	Price    int        `json:"price" gorm:"type:int"`
	Qty      int        `json:"qty" gorm:"type:int"`
	Image    string     `json:"image" gorm:"type:text"`
	SellerID int        `json:"seller_id"`
	Seller   UserRespon `json:"seller"`
}
