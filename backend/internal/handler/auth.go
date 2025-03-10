package handler

import (
	"net/http"

	"superbank/internal/model"
	"superbank/pkg/util"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleLogin(c *gin.Context) {
	var req model.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.NewResponse("Invalid request", nil, http.StatusBadRequest).ReturnGin(c)
		return
	}

	token, err := s.authService.Login(req)
	if err != nil {
		util.NewResponse("Invalid credentials", nil, http.StatusUnauthorized).ReturnGin(c)
		return
	}

	util.NewResponse("Login successful", model.TokenResponse{Token: token}, http.StatusOK).ReturnGin(c)
}