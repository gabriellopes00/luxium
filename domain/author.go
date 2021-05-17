package domain

import (
	"time"
	// uuid "github.com/satori/go.uuid"
	// "golang.org/x/crypto/bcrypt"
)

type Author struct {
	ID          string 		`json:"id"`
	Name        string 		`json:"name"`
	Password    string 		`json:"-"`
	Email       string 		`json:"email"`
	Instagram   string 		`json:"instagram"`
	Facebook    string 		`json:"facebook"`
	Twitter     string 		`json:"twitter"`
	Avatar      string 		`json:"avatar"`
	Description string 		`json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewAuthor(name string, email string, password string, description string, instagram string, 
facebook string, twitter string, avatar string) (*Author, error) {
	author := Author{
		Name: name,
		Email: email,
		Instagram: instagram,
		Facebook: facebook,
		Twitter: twitter,
		Avatar: avatar,
		Description: description,
	}

	// author.ID = uuid.NewV4().String()
	// password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	// if err != nil {
	// 	return err
	// }
	author.CreatedAt = time.Now().Local()
	author.UpdatedAt = time.Now().Local()
	return &author, nil
}