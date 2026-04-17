package repositories

import (
	"database/sql"
	"finalproject-it/backend/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(
		"SELECT id, username, password_hash, full_name, role, created_at FROM users WHERE username = $1",
		username,
	).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.FullName, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	user := &models.User{}
	err := r.db.QueryRow(
		"SELECT id, username, password_hash, full_name, role, created_at FROM users WHERE id = $1",
		id,
	).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.FullName, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, username, full_name, role, created_at FROM users ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.FullName, &u.Role, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.QueryRow(
		"INSERT INTO users (username, password_hash, full_name, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at",
		user.Username, user.PasswordHash, user.FullName, user.Role,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) Update(user *models.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET full_name = $1, role = $2 WHERE id = $3",
		user.FullName, user.Role, user.ID,
	)
	return err
}

func (r *UserRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
