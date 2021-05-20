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
	ID          string    `json:"id" gorm:"type:uuid;primary_key"`
	Name        string    `json:"name" gorm:"type:varchar(255)"`
	Password    string    `json:"-" gorm:"type:varchar(255);unique_index"`
	Email       string    `json:"email" gorm:"type:varchar(255);unique_index"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(255);unique_index"`
	Description string    `json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
	isNull := validator.IsNull(author.Name)
	if isNull {
		return errors.New("")
	}

	isNull = validator.IsNull(author.Password)
	if isNull {
		return errors.New("")
	}

	isEmail := validator.IsEmail(author.Email)
	if !isEmail {
		return errors.New("")

	}

	isUrl := validator.IsURL(author.Avatar)
	if !isUrl {
		return errors.New("")
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
