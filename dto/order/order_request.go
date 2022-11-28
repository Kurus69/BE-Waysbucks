package orderdto

type AddOrder struct {
	Qty       int   `json:"qty" form:"qty"`
	Price     int   `json:"price" form:"price"`
	ProductID int   `json:"product_id" form:"product_id"`
	TopingID  []int `json:"toping_id" form:"toping_id"`
	BuyyerID  int   `json:"buyyer_id" form:"buyyer_id"`
}

type UpdateOrder struct {
	Qty int `json:"qty" form:"qty"`
}
