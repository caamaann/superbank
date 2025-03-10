package service

import (
	"superbank/internal/model"
	"superbank/internal/repository/postgres"
)

type CustomerService interface {
	SearchCustomer(query string) (*model.Customer, error)
	GetCustomerByID(id string) (*model.Customer, error)
}

type customerService struct {
	repo postgres.CustomerRepository
}

func NewCustomerService(repo postgres.CustomerRepository) CustomerService {
	return &customerService{
		repo: repo,
	}
}

func (s *customerService) SearchCustomer(query string) (*model.Customer, error) {
	return s.repo.FindByQuery(query)
}

func (s *customerService) GetCustomerByID(id string) (*model.Customer, error) {
	return s.repo.GetByID(id)
}