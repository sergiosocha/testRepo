package repository

import (
	"database/sql"
	"miniApi_BRM/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetAll() ([]domain.User, error)
	GetByID(id int) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int) error
}

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func (r *MySQLUserRepository) GetAll() ([]domain.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *MySQLUserRepository) GetByID(id int) (*domain.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?`

	var user domain.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MySQLUserRepository) Update(user *domain.User) error {
	query := `UPDATE users SET name = ?, email = ?, updated_at = NOW() WHERE id = ?`
	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *MySQLUserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
