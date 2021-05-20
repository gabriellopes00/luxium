package main

import (
	"blog-golang/domain"
	database "blog-golang/infra"
	"blog-golang/infra/repositories"
	"fmt"
	"log"
)

func main() {
	author := domain.Author{
		Name:        "Gabriel Luís Lopes",
		Email:       "gabriellopes00@mail.com",
		Password:    "golang1234",
		Description: "Golang developer...",
		Avatar:      "https://avatars.githubusercontent.com/u/69465943?v=4",
	}

	db := database.ConnectPg()
	authorRepo := repositories.GormAuthorRepository{Db: db}

	newAuthor, err := authorRepo.Insert(&author)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newAuthor)
}
