package payment_service

import (
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
)

type PaymentService interface {
	CreatePayment(request request_schemas.CreatePaymentRequest, authHeader string) (*response_schemas.CreatePaymentResponse, error)
}
