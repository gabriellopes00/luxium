package main

import (
	"blog-golang/domain"
	"fmt"
	"log"
)

func main() {
	gabriel := domain.Author{
		Name:        "gabrielloes",
		Password:    "gabriels",
		Email:       "gabriel@mail.com",
		Avatar:      "https://picture.png",
		Description: "lorem ipsum dolor sit amet...",
	}
	err := gabriel.Create()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gabriel)
}
