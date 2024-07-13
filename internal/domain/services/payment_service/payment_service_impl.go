package payment_service

import (
	"SquarePosSystem/internal/domain/clients/square_client"
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	"log"
)

const paymentServiceLogPrefix = "payment_service_impl"

type paymentService struct {
	client square_client.PaymentClient
}

func NewPaymentService(client square_client.PaymentClient) PaymentService {
	return &paymentService{client: client}
}

func (p paymentService) CreatePayment(request request_schemas.CreatePaymentRequest, authHeader string) (*response_schemas.CreatePaymentResponse, error) {

	internalReq := request_schemas.CreatePaymentSquareRequest{
		SourceId:                   "CASH",
		AcceptPartialAuthorization: false,
		CashDetails: request_schemas.BuyerSuppliedMoney{
			request_schemas.Money{
				Amount:   request.BillAmount + request.TipAmount,
				Currency: "LKR",
			},
		},
		OrderId: request.OrderID,
		AmountMoney: request_schemas.Money{
			Amount:   request.BillAmount,
			Currency: "LKR",
		},
		TipMoney: request_schemas.Money{
			Amount:   request.TipAmount,
			Currency: "LKR",
		},
	}

	internalResp, err := p.client.CreatePayment(internalReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", paymentServiceLogPrefix, err)
		return nil, err
	}

	outoingResp := response_schemas.CreatePaymentResponse{
		PaymentId:   internalResp.Payment.Id,
		PaymentTime: internalResp.Payment.CreatedAt,
		Amount: response_schemas.Money{
			Amount:   internalResp.Payment.Amount.Amount,
			Currency: internalResp.Payment.Amount.Currency,
		},
		Tip: response_schemas.Money{
			Amount:   internalResp.Payment.Tip.Amount,
			Currency: internalResp.Payment.Tip.Currency,
		},
		Total: response_schemas.Money{
			Amount:   internalResp.Payment.Total.Amount,
			Currency: internalResp.Payment.Total.Currency,
		},
		ReceiptNumber: internalResp.Payment.ReceiptNumber,
		ReceiptUrl:    internalResp.Payment.ReceiptUrl,
	}

	return &outoingResp, nil
}
