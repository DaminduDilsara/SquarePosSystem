package response_schemas

import "time"

type CreateOrderSquareResponse struct {
	Order SquareOrder `json:"order"`
}

type CreateOrderResponse struct {
	OrderResponse
}
type SearchOrdersResponse struct {
	Orders []OrderResponse `json:"orders"`
}

type OrderResponse struct {
	Id       string    `json:"id"`
	OpenedAt time.Time `json:"opened_at"`
	IsClosed bool      `json:"is_closed"`
	Table    string    `json:"table"`
	Items    []Item    `json:"items"`
	Totals   Totals    `json:"totals"`
}

type Item struct {
	Name      string     `json:"name"`
	Comment   string     `json:"comment"`
	UnitPrice int        `json:"unit_price"`
	Quantity  int        `json:"quantity"`
	Discounts []Discount `json:"discounts"`
	Modifiers []Modifier `json:"modifiers"`
	Amount    int        `json:"amount"`
}

type Discount struct {
	Name         string `json:"name"`
	IsPercentage bool   `json:"is_percentage"`
	Value        int    `json:"value"`
	Amount       int    `json:"amount"`
}

type Modifier struct {
	Name      string `json:"name"`
	UnitPrice int    `json:"unit_price"`
	Quantity  int    `json:"quantity"`
	Amount    int    `json:"amount"`
}

type Totals struct {
	Discounts     int `json:"discounts"`
	Due           int `json:"due"`
	Tax           int `json:"tax"`
	ServiceCharge int `json:"service_charge"`
	Paid          int `json:"paid"`
	Tips          int `json:"tips"`
	Total         int `json:"total"`
}

type SearchOrdersSquareResponse struct {
	Orders []SquareOrder `json:"orders"`
}

type SquareOrder struct {
	Id                      string     `json:"id"`
	LocationId              string     `json:"location_id"`
	LineItems               []LineItem `json:"line_items"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at"`
	State                   string     `json:"state"`
	TotalMoney              Money      `json:"total_money"`
	TotalTaxMoney           Money      `json:"total_tax_money"`
	TotalDiscountMoney      Money      `json:"total_discount_money"`
	TotalTipMoney           Money      `json:"total_tip_money"`
	TotalServiceChargeMoney Money      `json:"total_service_charge_money"`
	NetAmounts              NetAmounts `json:"net_amounts"`
	Source                  Source     `json:"source"`
}

type Source struct {
	Name string `json:"name"`
}

type LineItem struct {
	Uid                     string `json:"uid"`
	Name                    string `json:"name"`
	Quantity                string `json:"quantity"`
	BasePriceMoney          Money  `json:"base_price_money"`
	Note                    string `json:"note"`
	GrossSalesMoney         Money  `json:"gross_sales_money"`
	TotalTaxMoney           Money  `json:"total_tax_money"`
	TotalDiscountMoney      Money  `json:"total_discount_money"`
	TotalMoney              Money  `json:"total_money"`
	ItemType                string `json:"item_type"`
	TotalServiceChargeMoney Money  `json:"total_service_charge_money"`
}

type Money struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type NetAmounts struct {
	TotalMoney         Money `json:"total_money"`
	TaxMoney           Money `json:"tax_money"`
	DiscountMoney      Money `json:"discount_money"`
	TipMoney           Money `json:"tip_money"`
	ServiceChargeMoney Money `json:"service_charge_money"`
}
