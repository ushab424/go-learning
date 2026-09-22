package repository

import (
	"database/sql"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int) (*User, error) {
	user := &User{}
	err := r.db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email string) (*User, *string, error) {
	user := &User{}
	var hash string
	err := r.db.QueryRow("SELECT id, name, email, age, password_hash FROM users WHERE email = $1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &hash)
	if err != nil {
		return nil, nil, err
	}
	return user, &hash, nil
}

func (r *UserRepository) Create(name, email, hash string, age int) (int, error) {
	var userID int
	err := r.db.QueryRow("INSERT INTO users (name, email, password_hash, age) VALUES ($1, $2, $3, $4) RETURNING id", name, email, hash, age).Scan(&userID)
	return userID, err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
