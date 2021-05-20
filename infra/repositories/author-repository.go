package repositories

import (
	"blog-golang/domain"
	"log"

	"github.com/jinzhu/gorm"
)

type AuthorRepository interface {
	Insert(author *domain.Author) (*domain.Author, error)
}

type GormAuthorRepository struct {
	Db *gorm.DB
}

func (repo GormAuthorRepository) Insert(author *domain.Author) (*domain.Author, error) {
	err := author.Create()
	if err != nil {
		log.Fatal("Error durign author creation")
		return author, err
	}

	err = repo.Db.Create(author).Error
	if err != nil {
		log.Fatal("Error durign author insertion")
		return author, err
	}

	return author, nil
}
