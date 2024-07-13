package response_schemas

import "time"

type CreateOrderSquareResponse struct {
	Order struct {
		Id            string    `json:"id"`
		LocationId    string    `json:"location_id"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		State         string    `json:"state"`
		Version       int       `json:"version"`
		ReferenceId   string    `json:"reference_id"`
		TotalTaxMoney struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"total_tax_money"`
		TotalDiscountMoney struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"total_discount_money"`
		TotalTipMoney struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"total_tip_money"`
		TotalMoney struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"total_money"`
		TotalServiceChargeMoney struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"total_service_charge_money"`
		NetAmounts struct {
			TotalMoney struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			} `json:"total_money"`
			TaxMoney struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			} `json:"tax_money"`
			DiscountMoney struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			} `json:"discount_money"`
			TipMoney struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			} `json:"tip_money"`
			ServiceChargeMoney struct {
				Amount   int    `json:"amount"`
				Currency string `json:"currency"`
			} `json:"service_charge_money"`
		} `json:"net_amounts"`
		Source struct {
			Name string `json:"name"`
		} `json:"source"`
		CustomerId        string `json:"customer_id"`
		TicketName        string `json:"ticket_name"`
		NetAmountDueMoney struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"net_amount_due_money"`
	} `json:"order"`
}

type CreateOrderResponse struct {
	OrderResponse
}

type OrderResponse struct {
	Id       string    `json:"id"`
	OpenedAt time.Time `json:"opened_at"`
	IsClosed bool      `json:"is_closed"`
	Table    string    `json:"table"`
	Items    []struct {
		Name      string `json:"name"`
		Comment   string `json:"comment"`
		UnitPrice int    `json:"unit_price"`
		Quantity  int    `json:"quantity"`
		Discounts []struct {
			Name         string `json:"name"`
			IsPercentage bool   `json:"is_percentage"`
			Value        int    `json:"value"`
			Amount       int    `json:"amount"`
		} `json:"discounts"`
		Modifiers []struct {
			Name      string `json:"name"`
			UnitPrice int    `json:"unit_price"`
			Quantity  int    `json:"quantity"`
			Amount    int    `json:"amount"`
		} `json:"modifiers"`
		Amount int `json:"amount"`
	} `json:"items"`
	Totals struct {
		Discounts     int `json:"discounts"`
		Due           int `json:"due"`
		Tax           int `json:"tax"`
		ServiceCharge int `json:"service_charge"`
		Paid          int `json:"paid"`
		Tips          int `json:"tips"`
		Total         int `json:"total"`
	} `json:"totals"`
}
