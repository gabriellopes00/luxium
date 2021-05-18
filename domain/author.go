package domain

import (
	"context"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type Author struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"-"`
	Avatar      string    `json:"avatar"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (author *Author) CreateAuthor() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(author.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	author.Password = string(hash)
	author.ID = uuid.NewV4().String()

	author.CreatedAt = time.Now().Local()
	author.UpdatedAt = time.Now().Local()
	return nil
}

type AuthorUsecase interface {
	GetByID(ctx context.Context, id int64) (Author, error)
	Add(context.Context, *Author) error
}

type AuthorRepository interface {
	FindByID(ctx context.Context, id int64) (Author, error)
	Store(ctx context.Context, a *Author) error
}
