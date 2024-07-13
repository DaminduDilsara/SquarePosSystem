package request_schemas

type CreateOrderIncomingRequest struct {
	LocationId  string      `json:"location_id" binding:"required"`
	CustomerId  string      `json:"customer_id" binding:"required"`
	ReferenceId string      `json:"reference_id" binding:"required"`
	TableId     string      `json:"table_id" binding:"required"`
	LineItems   []LineItems `json:"line_items" binding:"required"`
}

type LineItems struct {
	ItemId   string `json:"catalog_object_id" binding:"required"`
	ItemName string `json:"name" binding:"required"`
	Note     string `json:"note"`
	Quantity string `json:"quantity" binding:"required"`
}

type CreateOrderSquareRequest struct {
	IdempotencyKey string `json:"idempotency_key"`
	Order          Order  `json:"order"`
}

type Order struct {
	LocationId  string `json:"location_id"`
	CustomerId  string `json:"customer_id"`
	ReferenceId string `json:"reference_id"`
	TicketName  string `json:"ticket_name"`
	Source      struct {
		Name string `json:"name"`
	} `json:"source"`
	LineItems []struct {
		CatalogObjectId string `json:"catalog_object_id"`
		LineItems
	} `json:"line_items"`
}

type SearchOrdersIncomingRequest struct {
	LocationId string `json:"location_id" binding:"required"`
	TableNo    string `json:"table_no"`
}

type SearchOrdersSquareRequest struct {
	LocationIds   []string `json:"location_ids" binding:"required"`
	ReturnEntries bool     `json:"return_entries"`
	Query         struct {
		Filter struct {
			SourceFilter struct {
				SourceNames []string `json:"source_names"`
			} `json:"source_filter"`
		} `json:"filter"`
	} `json:"query"`
}
