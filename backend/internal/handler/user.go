package handler

import (
	"net/http"

	"superbank/internal/model"
	"superbank/internal/repository/postgres"
	"superbank/pkg/util"

	"github.com/gin-gonic/gin"
)


type AddUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8"`
}

func (s *Server) handleAddUser(c *gin.Context) {
	var req AddUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.NewResponse("Invalid request", err.Error(), http.StatusBadRequest).ReturnGin(c)
		return
	}
	
	hashedPassword, err := postgres.HashPassword(req.Password)
	if err != nil {
		util.NewResponse("Failed to process password", nil, http.StatusInternalServerError).ReturnGin(c)
		return
	}

	
	err = s.userService.CreateUser(model.CreateUserInput{
		Username: req.Username,
		Password: hashedPassword,
	})

	if err != nil {
		if err.Error() == "username already exists" {
			util.NewResponse("Username already exists", nil, http.StatusConflict).ReturnGin(c)
		} else {
			util.NewResponse("Failed to create user", err.Error(), http.StatusInternalServerError).ReturnGin(c)
		}
		return
	}

	util.NewResponse("User created successfully", gin.H{
		"username": req.Username,
	}, http.StatusCreated).ReturnGin(c)
}
