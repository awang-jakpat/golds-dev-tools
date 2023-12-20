package payment

type PaymentInfo struct {
	ID     string
	Status string
}

type Payment interface {
	Pay(amount float64) (PaymentInfo, error)
}
