package square_client

import (
	"SquarePosSystem/configurations"
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const locationClientLogPrefix = "square_location_client_impl"

type squareClient struct {
	config *configurations.Config
}

// NewSquareClient creates a new SquareClient
func NewSquareClient(config *configurations.Config) LocationClient {
	return &squareClient{
		config: config,
	}
}

func (s squareClient) CreateLocation(request request_schemas.CreateLocationSquareRequest, authHeader string) (*response_schemas.CreateLocationSquareResponse, error) {
	url := fmt.Sprintf("%v/locations", s.config.SquareConfig.BaseUrl)
	method := "POST"

	payload, err := json.Marshal(request)
	if err != nil {
		log.Printf("%v - Error marshalling JSON: %v", locationClientLogPrefix, err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("%v - Error: %v", locationClientLogPrefix, err)
		return nil, err
	}
	req.Header.Add("Square-Version", "2024-06-04")
	req.Header.Add("Authorization", fmt.Sprintf("%v", authHeader))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Printf("%v - Error: %v", locationClientLogPrefix, err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("%v - Error: %v", locationClientLogPrefix, err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		var errorResponse response_schemas.SquareErrorResponse
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			log.Printf("%v - Error: %v", locationClientLogPrefix, err)
			return nil, err
		}
		return nil, fmt.Errorf("square API error: %v", errorResponse.Errors)
	}

	var internalResp response_schemas.CreateLocationSquareResponse
	if err := json.Unmarshal(body, &internalResp); err != nil {
		log.Printf("%v - Error: %v", locationClientLogPrefix, err)
		return nil, err
	}

	return &internalResp, nil
}
