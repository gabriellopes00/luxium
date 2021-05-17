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
	ID          string 		`json:"id" gorm:"type:uuid;primary_key"`
	Name        string 		`json:"name" gorm:"type:varchar(255)"`
	Password    string 		`json:"-" gorm:"type:varchar(255);unique_index"`
	Email       string 		`json:"email" gorm:"type:varchar(255);unique_index"`
	Instagram   string 		`json:"instagram" gorm:"type:varchar(255);unique_index"`
	Facebook    string 		`json:"facebook" gorm:"type:varchar(255);unique_index"`
	Twitter     string 		`json:"twitter" gorm:"type:varchar(255);unique_index"`
	Avatar      string 		`json:"avatar" gorm:"type:varchar(255);unique_index"`
	Description string 		`json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:datetime"`
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
