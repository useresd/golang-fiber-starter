package repositories

import (
	"context"

	"github.com/useresd/golang-fiber-starter/internal/database"
	"github.com/useresd/golang-fiber-starter/internal/models"
)

type UserRepository interface {
	Store(context.Context, *models.User) error
	FindMany(context.Context) ([]*models.User, error)
}

type DefaultUserRepository struct{}

func NewDefaultUserRepository() *DefaultUserRepository {
	return &DefaultUserRepository{}
}

func (r *DefaultUserRepository) Store(ctx context.Context, user *models.User) error {

	tx, err := database.Tx(ctx)

	if err != nil {
		return err
	}

	if _, err := tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email); err != nil {
		return err
	}

	return nil
}

func (r *DefaultUserRepository) FindMany(ctx context.Context) ([]*models.User, error) {

	users := []*models.User{}

	tx, err := database.Tx(ctx)

	if err != nil {
		return users, err
	}

	rows, err := tx.QueryContext(ctx, "SELECT id, name, email FROM users")

	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}
