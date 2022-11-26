package topingdto

type AddToping struct {
	Title    string `json:"title" gorm:"type:text"`
	Image    string `json:"image" gorm:"type:text"`
	Price    int    `json:"price" gorm:"type:int"`
	SellerID int    `json:"seller_id"`
	Status   bool   `json:"status" gorm:"type:boolean"`
}
type UpdateToping struct {
	Title  string `json:"title" gorm:"type:text"`
	Price  int    `json:"price" gorm:"type:int"`
	Image  string `json:"image" gorm:"type:text"`
	Status bool   `json:"status" gorm:"type:boolean"`
}
