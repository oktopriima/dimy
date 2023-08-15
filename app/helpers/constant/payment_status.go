package constant

type PaymentStatus string

func (p PaymentStatus) String() string {
	return string(p)
}

const PaymentStatusPending PaymentStatus = "pending"
const PaymentStatusSuccess PaymentStatus = "success"
