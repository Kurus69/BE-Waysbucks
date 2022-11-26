package productdto

type AddProduct struct {
	Title    string `json:"title" form:"title" gorm:"type:text"`
	Price    int    `json:"price" form:"price" gorm:"type:int"`
	Qty      int    `json:"qty" form:"qty" gorm:"type:int"`
	Image    string `json:"image" form:"image" gorm:"type:text"`
	SellerID int    `json:"seller_id" gorm:"type:int"`
}

type UpdateProduct struct {
	Title string `json:"title" form:"title" gorm:"type:text"`
	Price int    `json:"price" form:"price" gorm:"type:int"`
	Qty   int    `json:"qty" form:"qty" gorm:"type:int"`
	Image string `json:"image" form:"image" gorm:"type:text"`
}
