package transactiondto

type Checkout struct {
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	CartID      []int  `json:"cart_id" form:"cart_id"`
	Total_price int    `json:"total_price" form:"total_price"`
	Status      string `json:"status" form:"status"`
	AccountID   int    `json:"account_id" form:"account_id"`
}
