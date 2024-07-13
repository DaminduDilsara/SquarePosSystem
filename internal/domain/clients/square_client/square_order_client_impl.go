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

const orderClientLogPrefix = "square_order_client_impl"

type squareOrderClient struct {
	config *configurations.Config
}

// NewSquareOrderClient creates a new SquareClient
func NewSquareOrderClient(config *configurations.Config) OrderClient {
	return &squareOrderClient{
		config: config,
	}
}

func (s squareOrderClient) CreateOrder(request request_schemas.CreateOrderSquareRequest, authHeader string) (*response_schemas.CreateOrderSquareResponse, error) {
	request.IdempotencyKey = uuid.NewString()

	url := fmt.Sprintf("%v/orders", s.config.SquareConfig.BaseUrl)
	method := "POST"

	payload, err := json.Marshal(request)
	if err != nil {
		log.Printf("%v - Error marshalling JSON: %v", orderClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Making request to %s with payload %s", orderClientLogPrefix, url, string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("%v - Error: %v", orderClientLogPrefix, err)
		return nil, err
	}
	req.Header.Add("Square-Version", s.config.SquareConfig.SquareVersion)
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", orderClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", orderClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Response status: %s, body: %s", orderClientLogPrefix, res.Status, string(body))

	if res.StatusCode != http.StatusOK {
		var errorResponse response_schemas.SquareErrorResponse
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", orderClientLogPrefix, err)
			return nil, err
		}
		log.Printf("%v - Square API error: %v", orderClientLogPrefix, errorResponse.Errors)
		return nil, fmt.Errorf("square API error: %v", errorResponse.Errors)
	}

	var internalResp response_schemas.CreateOrderSquareResponse
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", orderClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Successfully created order: %v", orderClientLogPrefix, internalResp)
	return &internalResp, nil
}

func (s *squareOrderClient) SearchOrders(request request_schemas.SearchOrdersSquareRequest, authHeader string) (*response_schemas.SearchOrdersSquareResponse, error) {

	url := fmt.Sprintf("%v/orders/search", s.config.SquareConfig.BaseUrl)
	method := "POST"

	payload, err := json.Marshal(request)
	if err != nil {
		log.Printf("%v - Error marshalling JSON: %v", orderClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Making request to %s with payload %s", orderClientLogPrefix, url, string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("%v - Error: %v", orderClientLogPrefix, err)
		return nil, err
	}
	req.Header.Add("Square-Version", s.config.SquareConfig.SquareVersion)
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", orderClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", orderClientLogPrefix, err)
		return nil, err
	}

	log.Printf("%v - Response status: %s, body: %s", orderClientLogPrefix, res.Status, string(body))

	if res.StatusCode != http.StatusOK {
		var errorResponse response_schemas.SquareErrorResponse
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", orderClientLogPrefix, err)
			return nil, err
		}
		log.Printf("%v - Square API error: %v", orderClientLogPrefix, errorResponse.Errors)
		return nil, fmt.Errorf("square API error: %v", errorResponse.Errors)
	}

	var internalResp response_schemas.SearchOrdersSquareResponse
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", orderClientLogPrefix, err)
		return nil, err
	}

	log.Printf("Successfully retrieved orders: %v", internalResp)
	return &internalResp, nil
}
