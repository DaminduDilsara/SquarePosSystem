package response_schemas

import "time"

type CreateLocationSquareResponse struct {
	Location struct {
		Id            string    `json:"id"`
		Name          string    `json:"name"`
		Timezone      string    `json:"timezone"`
		Capabilities  []string  `json:"capabilities"`
		Status        string    `json:"status"`
		CreatedAt     time.Time `json:"created_at"`
		MerchantId    string    `json:"merchant_id"`
		Country       string    `json:"country"`
		LanguageCode  string    `json:"language_code"`
		Currency      string    `json:"currency"`
		BusinessName  string    `json:"business_name"`
		Type          string    `json:"type"`
		BusinessEmail string    `json:"business_email"`
		Description   string    `json:"description"`
		Mcc           string    `json:"mcc"`
	} `json:"location"`
}

type CreateLocationResponse struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	Description   string    `json:"description"`
	BusinessEmail string    `json:"business_email"`
}
