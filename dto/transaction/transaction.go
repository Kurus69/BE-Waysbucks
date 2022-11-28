package transactiondto

type AddTransaction struct {
	Name         string `json:"name" form:"name"`
	Address      string `json:"address" form:"address"`
	Status       string `json:"status" form:"status"`
	Order        []int  `json:"order" form:"order"`
	AccountOrder int    `json:"account_order" form:"account_order"`
}
type UpdateTransaction struct {
	Status string `json:"status" form:"status"`
}
