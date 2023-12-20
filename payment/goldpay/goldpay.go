package goldpay

import (
	"net/http"

	"github.com/awang-jakpat/golds-dev-tools/httphelper"
	"github.com/awang-jakpat/golds-dev-tools/payment"
)

type GoldpayConfig struct {
	// API url for goldpay payment
	GoldpayApiUrl string
}

type goldpayApiResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type goldPay struct {
	config *GoldpayConfig
}

func NewGoldpayPayment(config *GoldpayConfig) payment.Payment {
	return &goldPay{
		config: config,
	}
}

// Pay will invoke goldpay API to make a transfer transaction from user to upstore wallet
func (gp *goldPay) Pay(amount float64) (payment.PaymentInfo, error) {
	helper := httphelper.HttpHelper{}

	httpReq := helper.Request(http.MethodPost, gp.config.GoldpayApiUrl, map[string]any{
		"weight": amount,
	})

	var response goldpayApiResponse
	if err := httpReq.Decode(&response); err != nil {
		return payment.PaymentInfo{}, err
	}

	return payment.PaymentInfo{
		ID:     response.ID,
		Status: response.Status,
	}, nil
}
