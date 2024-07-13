package v1

import (
	"SquarePosSystem/internal/domain/entities/schemas/response_schemas"
	"SquarePosSystem/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupRouter(controller *ControllerV1) *gin.Engine {
	r := gin.Default()
	v1Group := r.Group("/v1.0")
	{
		v1Group.POST("/location/create", controller.CreateLocationController)
		v1Group.POST("/order/create", controller.CreateOrderController)
		v1Group.POST("/orders/search", controller.SearchOrdersController)
		v1Group.POST("/orders/find", controller.FindOrdersController)
		v1Group.POST("/payment/create", controller.CreatePaymentController)
	}
	return r
}

func TestCreateLocationController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLocationService := mocks.NewMockLocationService(ctrl)

	controller := NewControllerV1(mockLocationService, nil, nil)

	router := setupRouter(controller)

	t.Run("success", func(t *testing.T) {
		requestBody := `{
						  "business_email": "business1@gmail.com",
						  "description": "Business1",
						  "business_name": "Business1"
						}`
		authHeader := "Bearer token"

		mockLocationService.EXPECT().
			CreateLocation(gomock.Any(), authHeader).
			Return(&response_schemas.CreateLocationResponse{Id: "location1"}, nil)

		req, _ := http.NewRequest("POST", "/v1.0/location/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"id":"location1"`)
	})

	t.Run("missing authorization header", func(t *testing.T) {
		requestBody := `{
						  "business_email": "business1@gmail.com",
						  "description": "Business1",
						  "business_name": "Business1"
						}`

		req, _ := http.NewRequest("POST", "/v1.0/location/create", strings.NewReader(requestBody))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"Authorization header is required"`)
	})

	t.Run("bad request", func(t *testing.T) {
		requestBody := `{"name":123}` // Invalid JSON
		authHeader := "Bearer token"

		req, _ := http.NewRequest("POST", "/v1.0/location/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error"`)
	})

	t.Run("internal server error", func(t *testing.T) {
		requestBody := `{
						  "business_email": "business1@gmail.com",
						  "description": "Business1",
						  "business_name": "Business1"
						}`
		authHeader := "Bearer token"

		mockLocationService.EXPECT().
			CreateLocation(gomock.Any(), authHeader).
			Return(nil, errors.New("internal error"))

		req, _ := http.NewRequest("POST", "/v1.0/location/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"internal error"`)
	})
}

func TestCreateOrderController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderService := mocks.NewMockOrderService(ctrl)

	controller := NewControllerV1(nil, mockOrderService, nil)

	router := setupRouter(controller)

	t.Run("success", func(t *testing.T) {
		requestBody := `{
							"location_id": "L3Q8ZJYHWS11N",
							"customer_id": "001",
							"reference_id": "001uiqydefiuvqe",
							"table_id": "002",
							"line_items": [
								{
									"catalog_object_id": "UXHZKPGKOKXRINRN2NNWYMYR",
									"name": "fish bun",
									"note": "toasted",
									"quantity": "5"
								}
							]
						}`
		authHeader := "Bearer token"

		mockOrderService.EXPECT().
			CreateOrder(gomock.Any(), authHeader).
			Return(&response_schemas.CreateOrderResponse{OrderResponse: response_schemas.OrderResponse{
				Id: "001",
			}}, nil)

		req, _ := http.NewRequest("POST", "/v1.0/order/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"id":"001"`)
	})

	t.Run("missing authorization header", func(t *testing.T) {
		requestBody := `{
							"location_id": "L3Q8ZJYHWS11N",
							"customer_id": "001",
							"reference_id": "001uiqydefiuvqe",
							"table_id": "002",
							"line_items": [
								{
									"catalog_object_id": "UXHZKPGKOKXRINRN2NNWYMYR",
									"name": "fish bun",
									"note": "toasted",
									"quantity": "5"
								}
							]
						}`

		req, _ := http.NewRequest("POST", "/v1.0/order/create", strings.NewReader(requestBody))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"Authorization header is required"`)
	})

	t.Run("bad request", func(t *testing.T) {
		requestBody := `{"name":123}` // Invalid JSON
		authHeader := "Bearer token"

		req, _ := http.NewRequest("POST", "/v1.0/order/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error"`)
	})

	t.Run("internal server error", func(t *testing.T) {
		requestBody := `{
							"location_id": "L3Q8ZJYHWS11N",
							"customer_id": "001",
							"reference_id": "001uiqydefiuvqe",
							"table_id": "002",
							"line_items": [
								{
									"catalog_object_id": "UXHZKPGKOKXRINRN2NNWYMYR",
									"name": "fish bun",
									"note": "toasted",
									"quantity": "5"
								}
							]
						}`
		authHeader := "Bearer token"

		mockOrderService.EXPECT().
			CreateOrder(gomock.Any(), authHeader).
			Return(nil, errors.New("internal error"))

		req, _ := http.NewRequest("POST", "/v1.0/order/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"internal error"`)
	})
}

func TestSearchOrdersController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderService := mocks.NewMockOrderService(ctrl)

	controller := NewControllerV1(nil, mockOrderService, nil)

	router := setupRouter(controller)

	t.Run("success", func(t *testing.T) {
		requestBody := `{
							"location_id": "L3Q8ZJYHWS11N",
							"table_no": "002"
						}`
		authHeader := "Bearer token"

		mockOrderService.EXPECT().
			SearchOrders(gomock.Any(), authHeader).
			Return(&response_schemas.SearchOrdersResponse{Orders: []response_schemas.OrderResponse{
				{
					Id: "001",
				},
			}}, nil)

		req, _ := http.NewRequest("POST", "/v1.0/orders/search", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"id":"001"`)
	})

	t.Run("missing authorization header", func(t *testing.T) {
		requestBody := `{
							"location_id": "L3Q8ZJYHWS11N",
							"table_no": "002"
						}`

		req, _ := http.NewRequest("POST", "/v1.0/orders/search", strings.NewReader(requestBody))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"Authorization header is required"`)
	})

	t.Run("bad request", func(t *testing.T) {
		requestBody := `{"name":123}` // Invalid JSON
		authHeader := "Bearer token"

		req, _ := http.NewRequest("POST", "/v1.0/orders/search", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error"`)
	})

	t.Run("internal server error", func(t *testing.T) {
		requestBody := `{
							"location_id": "L3Q8ZJYHWS11N",
							"table_no": "002"
						}`
		authHeader := "Bearer token"

		mockOrderService.EXPECT().
			SearchOrders(gomock.Any(), authHeader).
			Return(nil, errors.New("internal error"))

		req, _ := http.NewRequest("POST", "/v1.0/orders/search", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"internal error"`)
	})
}

func TestFindOrdersController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderService := mocks.NewMockOrderService(ctrl)

	controller := NewControllerV1(nil, mockOrderService, nil)

	router := setupRouter(controller)

	t.Run("success", func(t *testing.T) {
		requestBody := `{
							"order_ids": [
								"Z64rklvu3f4CRBHeBOggqXIdj4OZY"
							],
							"location_id": "L3Q8ZJYHWS11N"
						}`
		authHeader := "Bearer token"

		mockOrderService.EXPECT().
			FindOrders(gomock.Any(), authHeader).
			Return(&response_schemas.FindOrdersResponse{Orders: []response_schemas.OrderResponse{
				{
					Id: "001",
				},
			}}, nil)

		req, _ := http.NewRequest("POST", "/v1.0/orders/find", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"id":"001"`)
	})

	t.Run("missing authorization header", func(t *testing.T) {
		requestBody := `{
							"order_ids": [
								"Z64rklvu3f4CRBHeBOggqXIdj4OZY"
							],
							"location_id": "L3Q8ZJYHWS11N"
						}`

		req, _ := http.NewRequest("POST", "/v1.0/orders/find", strings.NewReader(requestBody))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"Authorization header is required"`)
	})

	t.Run("bad request", func(t *testing.T) {
		requestBody := `{"name":123}` // Invalid JSON
		authHeader := "Bearer token"

		req, _ := http.NewRequest("POST", "/v1.0/orders/find", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error"`)
	})

	t.Run("internal server error", func(t *testing.T) {
		requestBody := `{
							"order_ids": [
								"Z64rklvu3f4CRBHeBOggqXIdj4OZY"
							],
							"location_id": "L3Q8ZJYHWS11N"
						}`
		authHeader := "Bearer token"

		mockOrderService.EXPECT().
			FindOrders(gomock.Any(), authHeader).
			Return(nil, errors.New("internal error"))

		req, _ := http.NewRequest("POST", "/v1.0/orders/find", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"internal error"`)
	})
}

func TestCreatePaymentController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentService := mocks.NewMockPaymentService(ctrl)

	controller := NewControllerV1(nil, nil, mockPaymentService)

	router := setupRouter(controller)

	t.Run("success", func(t *testing.T) {
		requestBody := `{
							"order_id": "TrgENvxgUwAECYxGBjnIM3tRMyOZY",
							"bill_amount": 5000,
							"tip_amount": 10
						}`
		authHeader := "Bearer token"

		mockPaymentService.EXPECT().
			CreatePayment(gomock.Any(), authHeader).
			Return(&response_schemas.CreatePaymentResponse{
				PaymentId: "001",
			}, nil)

		req, _ := http.NewRequest("POST", "/v1.0/payment/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), `"payment_id":"001"`)
	})

	t.Run("missing authorization header", func(t *testing.T) {
		requestBody := `{
							"order_id": "TrgENvxgUwAECYxGBjnIM3tRMyOZY",
							"bill_amount": 5000,
							"tip_amount": 10
						}`

		req, _ := http.NewRequest("POST", "/v1.0/payment/create", strings.NewReader(requestBody))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"Authorization header is required"`)
	})

	t.Run("bad request", func(t *testing.T) {
		requestBody := `{"name":123}` // Invalid JSON
		authHeader := "Bearer token"

		req, _ := http.NewRequest("POST", "/v1.0/payment/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), `"error"`)
	})

	t.Run("internal server error", func(t *testing.T) {
		requestBody := `{
							"order_id": "TrgENvxgUwAECYxGBjnIM3tRMyOZY",
							"bill_amount": 5000,
							"tip_amount": 10
						}`
		authHeader := "Bearer token"

		mockPaymentService.EXPECT().
			CreatePayment(gomock.Any(), authHeader).
			Return(nil, errors.New("internal error"))

		req, _ := http.NewRequest("POST", "/v1.0/payment/create", strings.NewReader(requestBody))
		req.Header.Set("Authorization", authHeader)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"internal error"`)
	})
}
