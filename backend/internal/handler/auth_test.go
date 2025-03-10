package handler

import (
	"bytes"
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

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Login(req model.LoginRequest) (string, error) {
	args := m.Called(req)
	return args.String(0), args.Error(1)
}

func (m *MockAuthService) ValidateToken(token string) (string, string, error) {
	args := m.Called(token)
	return args.String(0), args.String(1), args.Error(2)
}

func TestHandleLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    interface{}
		setupAuth      func(*MockAuthService, model.LoginRequest)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid login credentials",
			requestBody: model.LoginRequest{
				Username:    "admin",
				Password: "password",
			},
			setupAuth: func(mockAuth *MockAuthService, req model.LoginRequest) {
				mockAuth.On("Login", req).Return("jwt-token-123", nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"message": "Login successful",
				"data": map[string]interface{}{
					"access_token": "jwt-token-123",
				},
				"status": float64(http.StatusOK),
			},
		},
		{
			name: "Invalid login credentials",
			requestBody: model.LoginRequest{
				Username:    "admin",
				Password: "wrongpassword",
			},
			setupAuth: func(mockAuth *MockAuthService, req model.LoginRequest) {
				mockAuth.On("Login", req).Return("", errors.New("invalid credentials"))
			},
			expectedStatus: http.StatusUnauthorized,
			expectedBody: map[string]interface{}{
				"message": "Invalid credentials",
				"data":    nil,
				"status":  float64(http.StatusUnauthorized),
			},
		},
		{
			name:        "Invalid request body",
			requestBody: "invalid json",
			setupAuth:   func(mockAuth *MockAuthService, req model.LoginRequest) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "Invalid request",
				"data":    nil,
				"status":  float64(http.StatusBadRequest),
			},
		},
		{
			name: "Missing required fields",
			requestBody: map[string]string{
				"username": "admin",
			},
			setupAuth:   func(mockAuth *MockAuthService, req model.LoginRequest) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "Invalid request",
				"data":    nil,
				"status":  float64(http.StatusBadRequest),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockAuthService := new(MockAuthService)
			
			server := &Server{
				authService: mockAuthService,
			}
			
			router := gin.New()
			router.POST("/api/auth/login", server.handleLogin)
			
			var requestBody []byte
			var err error
			
			switch v := tc.requestBody.(type) {
			case string:
				requestBody = []byte(v)
			default:
				requestBody, err = json.Marshal(tc.requestBody)
				assert.NoError(t, err)
			}
			
			req, err := http.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			
			if loginReq, ok := tc.requestBody.(model.LoginRequest); ok {
				tc.setupAuth(mockAuthService, loginReq)
			}
			
			recorder := httptest.NewRecorder()
			
			router.ServeHTTP(recorder, req)
			
			assert.Equal(t, tc.expectedStatus, recorder.Code)
			
			var response map[string]interface{}
			err = json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err)
			
			assert.Equal(t, tc.expectedBody, response)
			
			mockAuthService.AssertExpectations(t)
		})
	}
}