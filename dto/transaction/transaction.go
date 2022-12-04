package transactiondto

type AddTransaction struct {
	Name    string         `json:"name" form:"name"`
	Address string         `json:"address" form:"address"`
	Order   []OrderRequest `json:"order" form:"order"`
	UserID  int            `json:"user_id" form:"user_order"`
}
type UpdateTransaction struct {
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
}
