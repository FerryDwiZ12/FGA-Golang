package repository

import (
	"context"

	"github.com/FerryDwiZ12/FGA-Golang/tree/master/assigments/finalproject/internal/adapters/storage/postgres"
	"github.com/FerryDwiZ12/FGA-Golang/tree/master/assigments/finalproject/internal/core/domain"
)

type UserRepository struct {
	Db postgres.DBTX
}

func NewUserRepository(db postgres.DBTX) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (r UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	var id int
	query := `
		INSERT INTO users (username, email, password, age, profile_image_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.Db.QueryRowContext(ctx, query, user.Username,
		user.Email, user.Password, user.Age, user.ProfileImageURL).Scan(&id)
	if err != nil {
		return nil, err
	}

	u := &domain.User{
		ID:              id,
		Username:        user.Username,
		Email:           user.Email,
		Password:        user.Password,
		Age:             user.Age,
		ProfileImageURL: user.ProfileImageURL,
	}

	return u, nil
}

func (r UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, username, email, password, age, profile_image_url
		FROM users
		WHERE email = $1
	`
	err := r.Db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username,
		&user.Email, &user.Password, &user.Age, &user.ProfileImageURL)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r UserRepository) GetUserById(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User

	query := `
		SELECT id, username, email, password, age, profile_image_url
		FROM users
		WHERE id = $1
	`
	err := r.Db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username,
		&user.Email, &user.Password, &user.Age, &user.ProfileImageURL)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r UserRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, age = $3, profile_image_url = $4, updated_at=NOW()
		WHERE id = $5
		RETURNING id, username, email, age, profile_image_url
	`

	err := r.Db.QueryRowContext(ctx, query, user.Username, user.Email,
		user.Age, user.ProfileImageURL, user.ID).Scan(&user.ID, &user.Username,
		&user.Email, &user.Age, &user.ProfileImageURL)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`
	_, err := r.Db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
