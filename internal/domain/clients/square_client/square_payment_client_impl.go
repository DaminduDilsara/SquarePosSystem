package square_client

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
)

type PaymentClient interface {
	CreatePayment(request request_schemas.CreatePaymentSquareRequest, authHeader string) (*response_schemas.CreatePaymentSquareResponse, error)
}
