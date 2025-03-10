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

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(input model.CreateUserInput) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockUserService) GetUserByUsername(username string) (*model.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func TestHandleAddUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    interface{}
		setupMock      func(*MockUserService, model.CreateUserInput)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Successful user creation",
			requestBody: AddUserRequest{
				Username: "testuser",
				Password: "password123",
			},
			setupMock: func(mockService *MockUserService, input model.CreateUserInput) {
				mockService.On("CreateUser", mock.MatchedBy(func(i model.CreateUserInput) bool {
					return i.Username == "testuser" && i.Password != ""
				})).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"message": "User created successfully",
				"data": map[string]interface{}{
					"username": "testuser",
				},
				"status": float64(http.StatusCreated),
			},
		},
		{
			name: "Username already exists",
			requestBody: AddUserRequest{
				Username: "existinguser",
				Password: "password123",
			},
			setupMock: func(mockService *MockUserService, input model.CreateUserInput) {
				mockService.On("CreateUser", mock.MatchedBy(func(i model.CreateUserInput) bool {
					return i.Username == "existinguser" && i.Password != ""
				})).Return(errors.New("username already exists"))
			},
			expectedStatus: http.StatusConflict,
			expectedBody: map[string]interface{}{
				"message": "Username already exists",
				"data":    nil,
				"status":  float64(http.StatusConflict),
			},
		},
		{
			name: "Database error",
			requestBody: AddUserRequest{
				Username: "newuser",
				Password: "password123",
			},
			setupMock: func(mockService *MockUserService, input model.CreateUserInput) {
				mockService.On("CreateUser", mock.MatchedBy(func(i model.CreateUserInput) bool {
					return i.Username == "newuser" && i.Password != ""
				})).Return(errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"message": "Failed to create user",
				"data":    "database error",
				"status":  float64(http.StatusInternalServerError),
			},
		},
		{
			name: "Missing username",
			requestBody: map[string]interface{}{
				"password": "password123",
			},
			setupMock: func(mockService *MockUserService, input model.CreateUserInput) {
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "Invalid request",
				"status": float64(http.StatusBadRequest),
			},
		},
		{
			name: "Missing password",
			requestBody: map[string]interface{}{
				"username": "testuser",
			},
			setupMock: func(mockService *MockUserService, input model.CreateUserInput) {
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "Invalid request",
				"status": float64(http.StatusBadRequest),
			},
		},
		{
			name: "Username too short",
			requestBody: map[string]interface{}{
				"username": "ab",
				"password": "password123",
			},
			setupMock: func(mockService *MockUserService, input model.CreateUserInput) {
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "Invalid request",
				"status": float64(http.StatusBadRequest),
			},
		},
		{
			name: "Password too short",
			requestBody: map[string]interface{}{
				"username": "testuser",
				"password": "pass",
			},
			setupMock: func(mockService *MockUserService, input model.CreateUserInput) {
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"message": "Invalid request",
				"status": float64(http.StatusBadRequest),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockUserService)
			
			server := &Server{
				userService: mockService,
			}
			
			router := gin.New()
			router.POST("/users", server.handleAddUser)
			
			requestBody, err := json.Marshal(tc.requestBody)
			assert.NoError(t, err)
			
			req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(requestBody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			
			if _, ok := tc.requestBody.(AddUserRequest); ok {
				input := model.CreateUserInput{
					Username: tc.requestBody.(AddUserRequest).Username,
					Password: "",
				}
				tc.setupMock(mockService, input)
			}
			
			recorder := httptest.NewRecorder()
			
			router.ServeHTTP(recorder, req)
			
			assert.Equal(t, tc.expectedStatus, recorder.Code)
			
			var response map[string]interface{}
			err = json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err)
			
			if tc.expectedStatus == http.StatusBadRequest && response["data"] != nil {
				assert.Equal(t, tc.expectedBody["message"], response["message"])
				assert.Equal(t, tc.expectedBody["status"], response["status"])
			} else {
				assert.Equal(t, tc.expectedBody, response)
			}
			
			mockService.AssertExpectations(t)
		})
	}
}