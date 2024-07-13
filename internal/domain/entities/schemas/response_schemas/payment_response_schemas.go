package response_schemas

import "time"

type CreatePaymentSquareResponse struct {
	Payment struct {
		Id            string    `json:"id"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Amount        Money     `json:"amount_money"`
		Tip           Money     `json:"tip_money"`
		Status        string    `json:"status"`
		SourceType    string    `json:"source_type"`
		LocationId    string    `json:"location_id"`
		OrderId       string    `json:"order_id"`
		Total         Money     `json:"total_money"`
		ReceiptNumber string    `json:"receipt_number"`
		ReceiptUrl    string    `json:"receipt_url"`
	} `json:"payment"`
}

type CreatePaymentResponse struct {
	PaymentId     string    `json:"payment_id"`
	PaymentTime   time.Time `json:"payment_time"`
	Amount        Money     `json:"amount_money"`
	Tip           Money     `json:"tip_money"`
	Total         Money     `json:"total_money"`
	ReceiptNumber string    `json:"receipt_number"`
	ReceiptUrl    string    `json:"receipt_url"`
}
