package models

type Transaction struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name" gorm:"type:text"`
	Address string `json:"address" gorm:"type:text"`
	//====================================================================
	Status string     `json:"status" gorm:"type:text"`
	Total  int        `json:"total" gorm:"type:int"`
	Order  []Order    `json:"order" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID int        `json:"user_id"`
	User   UserRespon `json:"account" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
type TransactionResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
