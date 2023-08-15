package constant

type OrderStatus string

func (o OrderStatus) String() string {
	return string(o)
}

const OrderStatusPending OrderStatus = "pending"
const OrderStatusSuccess OrderStatus = "success"
