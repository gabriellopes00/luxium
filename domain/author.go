package domain

import (
	"fmt"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

func Entrypoint(){
	gabriel := Author{
		Name: "gabrielloes",
		Password: "gabriels",
		Instagram: "gabriel@mail.como",
		Email: "gabriel@mail.com",
		Facebook: "face@lopes.com",
		Twitter: "tt@lopes.com",
		Avatar: "https://picture.png",
		Description: "lorem ipsum dolor sit amet...",
	}

	err := gabriel.Create()
	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println(gabriel)
}

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

func (author *Author) Create() error {
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
