package handler

import (
	"database/sql"
	"log"
	"net/http"
	"superbank/pkg/util"

	"github.com/gin-gonic/gin"
)

func (s *Server) handleCustomerSearch(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		util.NewResponse("Search query is required",nil, http.StatusBadRequest).ReturnGin(c)
		return
	}

	customer, err := s.customerService.SearchCustomer(query)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NewResponse("Customer not found", nil, http.StatusNotFound).ReturnGin(c)
		} else {
			log.Printf("Error searching for customer: %v", err)
			util.NewResponse("Failed to search for customer", nil, http.StatusInternalServerError).ReturnGin(c)
		}
		return
	}

	util.NewResponse("Success to get customer", customer, http.StatusOK).ReturnGin(c)
}

func (s *Server) handleGetCustomer(c *gin.Context) {
	id := c.Param("id")

	customer, err := s.customerService.GetCustomerByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NewResponse("Customer not found", nil, http.StatusNotFound).ReturnGin(c)
		} else {
			log.Printf("Error retrieving customer: %v", err)
			util.NewResponse("Failed to retrieve customer", nil, http.StatusInternalServerError).ReturnGin(c)
		}
		return
	}

	util.NewResponse("Success to get customer", customer, http.StatusOK).ReturnGin(c)
}
