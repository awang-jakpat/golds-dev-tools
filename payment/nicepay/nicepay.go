package nicepay

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/awang-jakpat/golds-dev-tools/httphelper"
	"github.com/awang-jakpat/golds-dev-tools/payment"
	"github.com/awang-jakpat/golds-dev-tools/payment/paymentstatus"
)

type NicepayConfig struct {
	// API url for nicepay payment
	ImpAPIUrl string
	ImpKey    string
	ImpSecret string
	ImpUid    string
}

type nicepayPaymentAnnotation struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type nicepayApiResponse struct {
	Code     int                      `json:"code"`
	Message  string                   `json:"message"`
	Response nicepayPaymentAnnotation `json:"response"`
}

type nicepay struct {
	config *NicepayConfig
}

func NewNicepayPayment(config *NicepayConfig) payment.Payment {
	return &nicepay{
		config: config,
	}
}

// Pay will invoke port one API to check and validate if payment status is already paid or not
func (np *nicepay) Pay(amount float64) (payment.PaymentInfo, error) {
	accessToken, err := np.getAccessToken()
	if err != nil {
		return payment.PaymentInfo{}, err
	}

	helper := httphelper.NewHttpHelper(&http.Client{}, &httphelper.HttpConfig{})
	helper.SetHeaderFn(func(req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		return nil
	})

	httpReq := helper.Request(
		http.MethodGet,
		fmt.Sprintf("%s/payments/%s", np.config.ImpAPIUrl, np.config.ImpUid),
		map[string]any{},
	)

	var result nicepayApiResponse
	if err := httpReq.Decode(&result); err != nil {
		return payment.PaymentInfo{}, err
	}

	response := result.Response
	t, _ := json.Marshal(response)
	log.Println(t)

	var status string
	if response.Status == "paid" {
		status = paymentstatus.PAID
	}

	return payment.PaymentInfo{
		ID:     response.ID,
		Status: status,
	}, nil
}

func (np *nicepay) getAccessToken() (string, error) {
	type getAccessTokenResponse struct {
		Response struct {
			AccessToken string `json:"access_token"`
		} `json:"response"`
	}
	var response getAccessTokenResponse

	helper := httphelper.NewHttpHelper(&http.Client{}, &httphelper.HttpConfig{})
	helper.SetHeaderFn(func(req *http.Request) error {
		return nil
	})

	httpReq := helper.Request(http.MethodPost, np.config.ImpAPIUrl+"/users/getToken", map[string]any{
		"imp_key":    np.config.ImpKey,
		"imp_secret": np.config.ImpSecret,
	})

	if err := httpReq.Decode(&response); err != nil {
		return "", err
	}

	return response.Response.AccessToken, nil
}
