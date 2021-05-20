package domain

import (
	"errors"
	"log"
	"time"

	validator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type Author struct {
	ID          string    `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	Name        string    `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Password    string    `json:"-" gorm:"type:varchar(255);unique_index" valid:"notnull"`
	Email       string    `json:"email" gorm:"type:varchar(255);unique_index" valid:"notnull,email"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(255);unique_index" valid:"notnull,url"`
	Description string    `json:"description" gorm:"type:varchar(255)" valid:"notnull"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
}

func (author *Author) Create() error {
	err := author.validate()
	if err != nil {
		return err
	}

	err = author.build()
	if err != nil {
		return err
	}

	return nil
}

func (author *Author) validate() error {
	_, err := validator.ValidateStruct(author)
	if err != nil {
		return err
	}

	isValid := len(author.Name) >= 3
	if !isValid {
		return errors.New("invalid param: name")
	}

	isValid = len(author.Password) >= 4
	if !isValid {
		return errors.New("invalid param: password")
	}

	return nil
}

func (author *Author) build() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(author.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	author.Password = string(hash)
	author.ID = uuid.NewV4().String()

	author.CreatedAt = time.Now().Local()
	author.UpdatedAt = time.Now().Local()
	return nil
}
