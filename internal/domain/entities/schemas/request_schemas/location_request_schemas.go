package request_schemas

type CreateLocationIncomingRequest struct {
	BusinessEmail string `json:"business_email" binding:"required,email"`
	Description   string `json:"description" binding:"required"`
	BusinessName  string `json:"business_name" binding:"required"`
}

type CreateLocationSquareRequest struct {
	Location Location `json:"location"`
}

type Location struct {
	BusinessEmail string `json:"business_email"`
	Description   string `json:"description"`
	Name          string `json:"name"`
}
