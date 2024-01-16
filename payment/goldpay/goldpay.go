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
	Success  bool   `json:"success"`
	Messsage string `json:"message"`
	Data     struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	} `json:"data"`
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
	helper := httphelper.NewHttpHelper(&http.Client{}, &httphelper.HttpConfig{})

	helper.SetHeaderFn(func(req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", gp.config.AuthenticatedUserToken))
		return nil
	})
	httpReq := helper.Request(http.MethodPost, gp.config.GoldpayApiUrl, map[string]any{
		"weight": amount,
	})

	var response goldpayApiResponse
	if err := httpReq.Decode(&response); err != nil {
		return payment.PaymentInfo{}, err
	}

	var status string
	var data = response.Data
	if data.Status == "completed" {
		status = paymentstatus.PAID
	}

	return payment.PaymentInfo{
		ID:     data.ID,
		Status: status,
	}, nil
}
