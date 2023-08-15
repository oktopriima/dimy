package dto

type OrderCreateInput struct {
	CustomerID int `json:"customer_id"`
	AddressID  int `json:"address_id"`
	Products   []struct {
		ID       int     `json:"id"`
		Quantity float64 `json:"quantity"`
	} `json:"products"`
}

type OrderCreateResponse struct {
	OrderCreateInput
	TotalPrice float64 `json:"total_price"`
	OrderID    int     `json:"order_id"`
}
