package models

type Toping struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Title  string `json:"title" gorm:"type:text"`
	Image  string `json:"image" gorm:"type:text"`
	Price  int    `json:"price" gorm:"type:int"`
	Status bool   `json:"status" gorm:"type:boolean"`
}
