package order_service

import (
	"SquarePosSystem/internal/domain/clients/square_client"
	"SquarePosSystem/internal/domain/entities/schemas/request_schemas"
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	"log"
	"strconv"
)

const orderServiceLogPrefix = "order_service_impl"

// orderService implements the OrderService interface
type orderService struct {
	client square_client.OrderClient
}

// NewOrderService creates a new OrderService
func NewOrderService(client square_client.OrderClient) OrderService {
	return &orderService{client: client}
}

func (o orderService) CreateOrder(request request_schemas.CreateOrderIncomingRequest, authHeader string) (*response_schemas.CreateOrderResponse, error) {
	// Prepare the internal API request payload
	internalReq := request_schemas.CreateOrderSquareRequest{
		Order: request_schemas.Order{
			LocationId:  request.LocationId,
			CustomerId:  request.CustomerId,
			ReferenceId: request.ReferenceId,
			TicketName:  request.TableId,
			Source: struct {
				Name string `json:"name"`
			}{Name: request.TableId},
			LineItems: make([]struct {
				CatalogObjectId string `json:"catalog_object_id"`
				request_schemas.LineItems
			}, len(request.LineItems)),
		},
	}

	for i, item := range request.LineItems {
		internalReq.Order.LineItems[i] = struct {
			CatalogObjectId string `json:"catalog_object_id"`
			request_schemas.LineItems
		}{
			CatalogObjectId: item.ItemId,
			LineItems:       item,
		}
	}

	// Call the client function
	internalResp, err := o.client.CreateOrder(internalReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", orderServiceLogPrefix, err)
		return nil, err
	}

	response := o._convertToCreateOrderResponse(internalResp)

	return response, nil
}

func (o *orderService) SearchOrders(request request_schemas.SearchOrdersIncomingRequest, authHeader string) (*response_schemas.SearchOrdersResponse, error) {

	internalReq := request_schemas.SearchOrdersSquareRequest{
		LocationIds:   []string{request.LocationId},
		ReturnEntries: false,
	}

	if request.TableNo != "" {
		internalReq.Query.Filter.SourceFilter.SourceNames = []string{request.TableNo}
	}

	internalResp, err := o.client.SearchOrders(internalReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", orderServiceLogPrefix, err)
		return nil, err
	}

	response := o._convertToSearchOrdersSquareResponse(internalResp)

	return response, nil
}

func (o *orderService) FindOrders(request request_schemas.FindOrdersIncomingRequest, authHeader string) (*response_schemas.FindOrdersResponse, error) {
	internalReq := request_schemas.FindOrdersSquareRequest{
		OrderBatchRetrieveRequest: request_schemas.OrderBatchRetrieveRequest{
			OrderIds:   request.OrderIds,
			LocationId: request.LocationId,
		},
	}

	internalResp, err := o.client.FindOrders(internalReq, authHeader)
	if err != nil {
		log.Printf("%v - Error: %v", orderServiceLogPrefix, err)
		return nil, err
	}

	response := o._convertToFindOrdersSquareResponse(internalResp)

	return response, nil
}

func (o orderService) _convertToCreateOrderResponse(squareResp *response_schemas.CreateOrderSquareResponse) *response_schemas.CreateOrderResponse {
	order := o._convertSquareOrderToOrderResponse(squareResp.Order)

	resp := response_schemas.CreateOrderResponse{order}

	return &resp
}

func (o orderService) _convertSquareOrderToOrderResponse(squareOrder response_schemas.SquareOrder) response_schemas.OrderResponse {
	items := make([]response_schemas.Item, len(squareOrder.LineItems))
	for i, lineItem := range squareOrder.LineItems {
		items[i] = response_schemas.Item{
			Name:      lineItem.Name,
			Comment:   lineItem.Note,
			UnitPrice: lineItem.BasePriceMoney.Amount,
			Quantity:  o._stringToInt(lineItem.Quantity),
			Discounts: []response_schemas.Discount{},
			Modifiers: []response_schemas.Modifier{},
			Amount:    lineItem.TotalMoney.Amount,
		}
	}
	return response_schemas.OrderResponse{
		Id:       squareOrder.Id,
		OpenedAt: squareOrder.CreatedAt,
		IsClosed: squareOrder.State == "COMPLETED",
		Table:    squareOrder.Source.Name,
		Items:    items,
		Totals: response_schemas.Totals{
			Discounts:     squareOrder.TotalDiscountMoney.Amount,
			Due:           0,
			Tax:           squareOrder.TotalTaxMoney.Amount,
			ServiceCharge: squareOrder.TotalServiceChargeMoney.Amount,
			Paid:          0,
			Tips:          squareOrder.TotalTipMoney.Amount,
			Total:         squareOrder.TotalMoney.Amount,
		},
	}
}

func (o orderService) _stringToInt(s string) int {
	val, _ := strconv.Atoi(s)
	return val
}

func (o orderService) _convertToSearchOrdersSquareResponse(squareResponse *response_schemas.SearchOrdersSquareResponse) *response_schemas.SearchOrdersResponse {
	orders := make([]response_schemas.OrderResponse, len(squareResponse.Orders))
	for i, squareOrder := range squareResponse.Orders {
		orders[i] = o._convertSquareOrderToOrderResponse(squareOrder)
	}
	return &response_schemas.SearchOrdersResponse{
		Orders: orders,
	}
}

func (o orderService) _convertToFindOrdersSquareResponse(squareResponse *response_schemas.FindOrdersSquareResponse) *response_schemas.FindOrdersResponse {
	orders := make([]response_schemas.OrderResponse, len(squareResponse.Orders))
	for i, squareOrder := range squareResponse.Orders {
		orders[i] = o._convertSquareOrderToOrderResponse(squareOrder)
	}
	return &response_schemas.FindOrdersResponse{
		Orders: orders,
	}
}
