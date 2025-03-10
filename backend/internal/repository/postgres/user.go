package postgres

import (
	"database/sql"
	"errors"

	"superbank/internal/model"

	"golang.org/x/crypto/bcrypt"
)


type UserRepository interface {
	GetUserByUsername(username string) (*model.User, error)
	ValidateCredentials(username, password string) (*model.User, error)
	CreateUser(username, hashedPassword string) error
}

type userRepository struct {
	db *sql.DB
}


func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}


func (r *userRepository) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}

	err := r.db.QueryRow(`
		SELECT id, username, password
		FROM users
		WHERE username = $1
	`, username).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}


func (r *userRepository) ValidateCredentials(username, password string) (*model.User, error) {
	user, err := r.GetUserByUsername(username)
	if err != nil {
		
		return nil, errors.New("invalid credentials")
	}

	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}


func (r *userRepository) CreateUser(username, hashedPassword string) error {
	
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("username already exists")
	}

	
	_, err = r.db.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		username, hashedPassword,
	)

	return err
}



func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
