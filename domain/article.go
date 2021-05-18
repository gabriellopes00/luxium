package domain

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Article struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	Title     string    `json:"name" gorm:"type:varchar(255)"`
	Content   string    `json:"email" gorm:"type:text"`
	Author    string    `json:"author" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

func (article *Article) CreateArticle() error {
	article.ID = uuid.NewV4().String()

	article.CreatedAt = time.Now().Local()
	article.UpdatedAt = time.Now().Local()
	return nil
}

type ArticleUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Article, string, error)
	GetByID(ctx context.Context, id int64) (Article, error)
	GetByTitle(ctx context.Context, title string) (Article, error)
	Add(context.Context, *Article) error
	Delete(ctx context.Context, id int64) error
}

type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Article, nextCursor string, err error)
	FindByID(ctx context.Context, id int64) (Article, error)
	FindByTitle(ctx context.Context, title string) (Article, error)
	Store(ctx context.Context, a *Article) error
	Delete(ctx context.Context, id int64) error
}
