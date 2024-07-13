package request_schemas

type CreatePaymentSquareRequest struct {
	IdempotencyKey             string             `json:"idempotency_key"`
	SourceId                   string             `json:"source_id"`
	AcceptPartialAuthorization bool               `json:"accept_partial_authorization"`
	CashDetails                BuyerSuppliedMoney `json:"cash_details"`
	OrderId                    string             `json:"order_id"`
	AmountMoney                Money              `json:"amount_money"`
	TipMoney                   Money              `json:"tip_money"`
}

type BuyerSuppliedMoney struct {
	BuyerSuppliedMoney Money `json:"buyer_supplied_money"`
}

type CreatePaymentRequest struct {
	BillAmount int    `json:"bill_amount"  binding:"required"`
	TipAmount  int    `json:"tip_amount"`
	OrderID    string `json:"order_id"  binding:"required"`
}

type Money struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
