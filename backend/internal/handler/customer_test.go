package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"superbank/internal/model"
)

type MockCustomerService struct {
	mock.Mock
}

func (m *MockCustomerService) SearchCustomer(query string) (*model.Customer, error) {
	args := m.Called(query)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*model.Customer), args.Error(1)
}

func (m *MockCustomerService) GetCustomerByID(id string) (*model.Customer, error) {
	args := m.Called(id)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*model.Customer), args.Error(1)
}

var testCustomer = &model.Customer{
	ID:      "cust123",
	Name:    "John Doe",
	Email:   "john.doe@example.com",
	Phone:   "123-456-7890",
	Address: "123 Main St",
	BankAccounts: nil,
	Pockets: nil,
	TermDeposits: nil,
}

func TestHandleCustomerSearch(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		queryParam     string
		setupMock      func(*MockCustomerService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:       "Successful search",
			queryParam: "john",
			setupMock: func(mockService *MockCustomerService) {
				mockService.On("SearchCustomer", "john").Return(testCustomer, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "Success to get customer",
				"data": map[string]interface{}{
					"id":      "cust123",
					"name":    "John Doe",
					"email":   "john.doe@example.com",
					"phone":   "123-456-7890",
					"address": "123 Main St",
					"createdAt": "0001-01-01T00:00:00Z",
					"bankAccounts": interface{}(nil),
					"pockets": interface{}(nil),
					"termDeposits": interface{}(nil),
				},
				"status": float64(http.StatusOK),
			},
		},
		{
			name:       "Customer not found",
			queryParam: "nonexistent",
			setupMock: func(mockService *MockCustomerService) {
				mockService.On("SearchCustomer", "nonexistent").Return(nil, sql.ErrNoRows)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"message": "Customer not found",
				"data":    nil,
				"status":  float64(http.StatusNotFound),
			},
		},
		{
			name:       "Database error",
			queryParam: "error",
			setupMock: func(mockService *MockCustomerService) {
				mockService.On("SearchCustomer", "error").Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"message": "Failed to search for customer",
				"data":    nil,
				"status":  float64(http.StatusInternalServerError),
			},
		},
		{
			name:       "Empty query parameter",
			queryParam: "",
			setupMock:  func(mockService *MockCustomerService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "Search query is required",
				"data":    nil,
				"status":  float64(http.StatusBadRequest),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockCustomerService)
			
			tc.setupMock(mockService)
			
			server := &Server{
				customerService: mockService,
			}
			
			router := gin.New()
			router.GET("/api/customers/search", server.handleCustomerSearch)
			
			req, err := http.NewRequest(http.MethodGet, "/api/customers/search", nil)
			assert.NoError(t, err)
			
			if tc.queryParam != "" {
				q := req.URL.Query()
				q.Add("q", tc.queryParam)
				req.URL.RawQuery = q.Encode()
			}
			
			recorder := httptest.NewRecorder()
			
			router.ServeHTTP(recorder, req)
			
			assert.Equal(t, tc.expectedStatus, recorder.Code)
			
			var response map[string]interface{}
			err = json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err)
			
			assert.Equal(t, tc.expectedBody, response)
			
			mockService.AssertExpectations(t)
		})
	}
}

func TestHandleGetCustomer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		customerID     string
		setupMock      func(*MockCustomerService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:       "Successful retrieval",
			customerID: "cust123",
			setupMock: func(mockService *MockCustomerService) {
				mockService.On("GetCustomerByID", "cust123").Return(testCustomer, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "Success to get customer",
				"data": map[string]interface{}{
					"id":      "cust123",
					"name":    "John Doe",
					"email":   "john.doe@example.com",
					"phone":   "123-456-7890",
					"address": "123 Main St",
					"createdAt": "0001-01-01T00:00:00Z",
					"bankAccounts": interface{}(nil),
					"pockets": interface{}(nil),
					"termDeposits": interface{}(nil),
				},
				"status": float64(http.StatusOK),
			},
		},
		{
			name:       "Customer not found",
			customerID: "nonexistent",
			setupMock: func(mockService *MockCustomerService) {
				mockService.On("GetCustomerByID", "nonexistent").Return(nil, sql.ErrNoRows)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"message": "Customer not found",
				"data":    nil,
				"status":  float64(http.StatusNotFound),
			},
		},
		{
			name:       "Database error",
			customerID: "error",
			setupMock: func(mockService *MockCustomerService) {
				mockService.On("GetCustomerByID", "error").Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"message": "Failed to retrieve customer",
				"data":    nil,
				"status":  float64(http.StatusInternalServerError),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockCustomerService)
			
			tc.setupMock(mockService)
			
			server := &Server{
				customerService: mockService,
			}
			
			router := gin.New()
			router.GET("/api/customers/:id", server.handleGetCustomer)
			
			req, err := http.NewRequest(http.MethodGet, "/api/customers/"+tc.customerID, nil)
			assert.NoError(t, err)
			
			recorder := httptest.NewRecorder()
			
			router.ServeHTTP(recorder, req)
			
			assert.Equal(t, tc.expectedStatus, recorder.Code)
			
			var response map[string]interface{}
			err = json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err)
			
			assert.Equal(t, tc.expectedBody, response)
			
			mockService.AssertExpectations(t)
		})
	}
}