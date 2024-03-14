package payment

type PaymentInfo struct {
	ID      string
	Status  string
	Nicepay NicepayPaymentAnnotation
}

type Payment interface {
	Pay(amount float64) (PaymentInfo, error)
}

type NicepayPaymentAnnotation struct {
	ID                string   `json:"id"`
	ImpUID            string   `json:"imp_uid"`                       // Port One transaction identification number
	MerchantUID       string   `json:"merchant_uid"`                  // Merchant order number for payment
	PayMethod         string   `json:"pay_method,omitempty"`          // Payment method classification code (Optional)
	Channel           string   `json:"channel,omitempty"`             // Payment environment classification code (Optional)
	PGProvider        string   `json:"pg_provider,omitempty"`         // PG company classification code for payment (Optional)
	EmbPGProvider     string   `json:"emb_pg_provider,omitempty"`     // Hub type payment PG company classification code (Optional)
	PGTID             string   `json:"pg_tid,omitempty"`              // PG company transaction number for payment (Optional)
	PGID              string   `json:"pg_id,omitempty"`               // PG company store ID of payment request (Optional)
	Escrow            bool     `json:"escrow,omitempty"`              // Escrow payment status (Optional)
	ApplyNum          string   `json:"apply_num,omitempty"`           // Approval number (Optional)
	BankCode          string   `json:"bank_code,omitempty"`           // Bank standard code for payment (Optional)
	BankName          string   `json:"bank_name,omitempty"`           // Bank name for payment (Optional)
	CardCode          string   `json:"card_code,omitempty"`           // Card company code number for payment (Optional)
	CardName          string   `json:"card_name,omitempty"`           // Card company name for payment (Optional)
	CardIssuerCode    string   `json:"card_issuer_code,omitempty"`    // Card issuer code (Optional)
	CardIssuerName    string   `json:"card_issuer_name,omitempty"`    // Card issuer name (Optional)
	CardPublisherCode string   `json:"card_publisher_code,omitempty"` // Card publisher code number (Optional)
	CardPublisherName string   `json:"card_publisher_name,omitempty"` // Card publisher name (Optional)
	CardQuota         int      `json:"card_quota,omitempty"`          // Number of installment months (Optional)
	CardNumber        string   `json:"card_number,omitempty"`         // Card number (Optional)
	CardType          int      `json:"card_type,omitempty"`           // Card identification code (Optional)
	VBankCode         string   `json:"vbank_code,omitempty"`          // Virtual account bank standard code (Optional)
	VBankName         string   `json:"vbank_name,omitempty"`          // Virtual account bank name (Optional)
	VBankNum          string   `json:"vbank_num,omitempty"`           // Virtual account account number (Optional)
	VBankHolder       string   `json:"vbank_holder,omitempty"`        // Virtual account depositor (Optional)
	VBankDate         int      `json:"vbank_date,omitempty"`          // Virtual account deposit deadline (Optional)
	VBankIssuedAt     int      `json:"vbank_issued_at,omitempty"`     // Virtual account creation time (Optional)
	Name              string   `json:"name,omitempty"`                // Product name of payment (Optional)
	Amount            float64  `json:"amount"`                        // Payment amount of payment case
	CancelAmount      float64  `json:"cancel_amount"`                 // Cumulative cancellation amount of payments
	Currency          string   `json:"currency"`                      // Payment currency classification code
	BuyerName         string   `json:"buyer_name,omitempty"`          // Orderer name (Optional)
	BuyerEmail        string   `json:"buyer_email,omitempty"`         // Orderer email address (Optional)
	BuyerTel          string   `json:"buyer_tel,omitempty"`           // Orderer phone number (Optional)
	BuyerAddr         string   `json:"buyer_addr,omitempty"`          // Orderer's address (Optional)
	BuyerPostcode     string   `json:"buyer_postcode,omitempty"`      // Orderer zip code (Optional)
	CustomData        string   `json:"custom_data,omitempty"`         // Additional information (Optional)
	UserAgent         string   `json:"user_agent,omitempty"`          // UserAgent string of the terminal (Optional)
	Status            string   `json:"status"`                        // Payment status
	StartedAt         int      `json:"started_at,omitempty"`          // Request time (Optional)
	PaidAt            int      `json:"paid_at,omitempty"`             // Payment time (Optional)
	FailedAt          int      `json:"failed_at,omitempty"`           // Time of failure (Optional)
	CancelledAt       int      `json:"cancelled_at,omitempty"`        // Cancellation time (Optional)
	FailReason        string   `json:"fail_reason,omitempty"`         // Reason for payment failure (Optional)
	CancelReason      string   `json:"cancel_reason,omitempty"`       // Reason for payment cancellation (Optional)
	ReceiptURL        string   `json:"receipt_url,omitempty"`         // Sales slip URL (Optional)
	CancelReceiptURLs []string `json:"cancel_receipt_urls,omitempty"` // (Deprecated) Cancellation sales slip confirmation URL (Optional)
	CashReceiptIssued bool     `json:"cash_receipt_issued,omitempty"` // Issuance of cash receipt? (Optional)
	CustomerUID       string   `json:"customer_uid,omitempty"`        // Unique number identifying the buyer’s payment method (Optional)
	CustomerUIDUsage  string   `json:"customer_uid_usage,omitempty"`  // Identification code for buyer’s payment method using unique number (Optional)
}
