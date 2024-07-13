package square_client

import (
	"SquarePosSystem/configurations"
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
)

const paymentClientLogPrefix = "square_payment_client_impl"

type squarePaymentClient struct {
	config *configurations.Config
}

// NewSquarePaymentClient creates a new SquareClient
func NewSquarePaymentClient(config *configurations.Config) PaymentClient {
	return &squarePaymentClient{
		config: config,
	}
}

func (s squarePaymentClient) CreatePayment(request request_schemas.CreatePaymentSquareRequest, authHeader string) (*response_schemas.CreatePaymentSquareResponse, error) {

	request.IdempotencyKey = uuid.NewString()

	url := fmt.Sprintf("%v/payments", s.config.SquareConfig.BaseUrl)
	method := "POST"

	payload, err := json.Marshal(request)
	if err != nil {
		log.Printf("%v - Error marshalling JSON: %v", paymentClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Making request to %s with payload %s", paymentClientLogPrefix, url, string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("%v - Error: %v", paymentClientLogPrefix, err)
		return nil, err
	}
	req.Header.Add("Square-Version", s.config.SquareConfig.SquareVersion)
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", paymentClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", paymentClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Response status: %s, body: %s", paymentClientLogPrefix, res.Status, string(body))

	if res.StatusCode != http.StatusOK {
		var errorResponse response_schemas.SquareErrorResponse
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", paymentClientLogPrefix, err)
			return nil, err
		}
		log.Printf("%v - Square API error: %v", paymentClientLogPrefix, errorResponse.Errors)
		return nil, fmt.Errorf("square API error: %v", errorResponse.Errors)
	}

	var internalResp response_schemas.CreatePaymentSquareResponse
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", paymentClientLogPrefix, err)
		return nil, err
	}

	log.Printf("Successfully retrieved orders: %v", internalResp)
	return &internalResp, nil
}
