package postgres

import (
	"database/sql"
	"fmt"

	"superbank/internal/model"
)

type CustomerRepository interface {
	FindByQuery(query string) (*model.Customer, error)
	GetByID(id string) (*model.Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindByQuery(query string) (*model.Customer, error) {
	
	customer, err := r.GetByID(query)
	if err == nil {
		return customer, nil
	}

	
	var customerID string
	err = r.db.QueryRow(`
		SELECT id FROM customers 
		WHERE name ILIKE $1 OR email ILIKE $1 
		LIMIT 1
	`, "%"+query+"%").Scan(&customerID)

	if err != nil {
		return nil, err
	}

	
	return r.GetByID(customerID)
}

func (r *customerRepository) GetByID(id string) (*model.Customer, error) {
	customer := &model.Customer{}
	
	
	var profileImage sql.NullString
	err := r.db.QueryRow(`
		SELECT id, name, email, phone, address, profile_image, created_at 
		FROM customers 
		WHERE id = $1
	`, id).Scan(
		&customer.ID,
		&customer.Name,
		&customer.Email,
		&customer.Phone,
		&customer.Address,
		&profileImage,
		&customer.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	
	if profileImage.Valid {
		customer.ProfileImage = profileImage.String
	}

	
	bankAccounts, err := r.getBankAccounts(id)
	if err != nil {
		return nil, fmt.Errorf("error getting bank accounts: %w", err)
	}
	customer.BankAccounts = bankAccounts

	
	pockets, err := r.getPockets(id)
	if err != nil {
		return nil, fmt.Errorf("error getting pockets: %w", err)
	}
	customer.Pockets = pockets

	
	termDeposits, err := r.getTermDeposits(id)
	if err != nil {
		return nil, fmt.Errorf("error getting term deposits: %w", err)
	}
	customer.TermDeposits = termDeposits

	return customer, nil
}

func (r *customerRepository) getBankAccounts(customerID string) ([]model.BankAccount, error) {
	rows, err := r.db.Query(`
		SELECT id, account_number, account_type, balance, currency, is_active, created_at 
		FROM bank_accounts 
		WHERE customer_id = $1
	`, customerID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []model.BankAccount
	for rows.Next() {
		var account model.BankAccount
		err := rows.Scan(
			&account.ID,
			&account.AccountNumber,
			&account.AccountType,
			&account.Balance,
			&account.Currency,
			&account.IsActive,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *customerRepository) getPockets(customerID string) ([]model.Pocket, error) {
	rows, err := r.db.Query(`
		SELECT id, name, balance, currency, goal, created_at 
		FROM pockets 
		WHERE customer_id = $1
	`, customerID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pockets []model.Pocket
	for rows.Next() {
		var pocket model.Pocket
		err := rows.Scan(
			&pocket.ID,
			&pocket.Name,
			&pocket.Balance,
			&pocket.Currency,
			&pocket.Goal,
			&pocket.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		pockets = append(pockets, pocket)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return pockets, nil
}

func (r *customerRepository) getTermDeposits(customerID string) ([]model.TermDeposit, error) {
	rows, err := r.db.Query(`
		SELECT id, amount, currency, interest_rate, start_date, maturity_date, is_active 
		FROM term_deposits 
		WHERE customer_id = $1
	`, customerID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deposits []model.TermDeposit
	for rows.Next() {
		var deposit model.TermDeposit
		err := rows.Scan(
			&deposit.ID,
			&deposit.Amount,
			&deposit.Currency,
			&deposit.InterestRate,
			&deposit.StartDate,
			&deposit.MaturityDate,
			&deposit.IsActive,
		)
		if err != nil {
			return nil, err
		}
		deposits = append(deposits, deposit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return deposits, nil
}
