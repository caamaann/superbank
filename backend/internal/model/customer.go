package model

import (
	"time"
)

type Customer struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	Address      string         `json:"address"`
	ProfileImage string         `json:"profileImage,omitempty"`
	CreatedAt    time.Time      `json:"createdAt"`
	BankAccounts []BankAccount  `json:"bankAccounts"`
	Pockets      []Pocket       `json:"pockets"`
	TermDeposits []TermDeposit  `json:"termDeposits"`
}

type BankAccount struct {
	ID            string    `json:"id"`
	AccountNumber string    `json:"accountNumber"`
	AccountType   string    `json:"accountType"`
	Balance       float64   `json:"balance"`
	Currency      string    `json:"currency"`
	IsActive      bool      `json:"isActive"`
	CreatedAt     time.Time `json:"createdAt"`
}

type Pocket struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Goal      *float64  `json:"goal,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type TermDeposit struct {
	ID           string    `json:"id"`
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currency"`
	InterestRate float64   `json:"interestRate"`
	StartDate    time.Time `json:"startDate"`
	MaturityDate time.Time `json:"maturityDate"`
	IsActive     bool      `json:"isActive"`
}