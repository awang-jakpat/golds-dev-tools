package goldpay

import (
	"fmt"
	"net/http"

	"github.com/awang-jakpat/golds-dev-tools/httphelper"
	"github.com/awang-jakpat/golds-dev-tools/payment"
	"github.com/awang-jakpat/golds-dev-tools/payment/paymentstatus"
)

type GoldpayConfig struct {
	// API url for goldpay payment
	GoldpayApiUrl string

	// Authenticated user token
	AuthenticatedUserToken string
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

	httpReq.SetHeaderFn(func(req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gp.config.AuthenticatedUserToken))
		return nil
	})

	var response goldpayApiResponse
	if err := httpReq.Decode(&response); err != nil {
		return payment.PaymentInfo{}, err
	}

	var status string
	if response.Status == "completed" {
		status = paymentstatus.PAID
	}

	return payment.PaymentInfo{
		ID:     response.ID,
		Status: status,
	}, nil
}
