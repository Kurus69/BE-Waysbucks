package models

type Transaction struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Name        string `json:"name" gorm:"type:text"`
	Address     string `json:"address" gorm:"type:text"`
	Total_price int    `json:"total_price" gorm:"type:int"`
	Status      string `json:"status" gorm:"type:text"`

	//-------------------------------------------------------------//
	CartID int        `json:"cart_id"`
	Cart   CartRespon `json:"cart"`

	//=============================================================//
	AccountID int        `json:"account_id"`
	Account   UserRespon `json:"account"`
}
