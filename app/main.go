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
		Instagram:   "gabriel@mail.como",
		Email:       "gabriel@mail.com",
		Facebook:    "face@lopes.com",
		Twitter:     "tt@lopes.com",
		Avatar:      "https://picture.png",
		Description: "lorem ipsum dolor sit amet...",
	}
	err := gabriel.Create()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gabriel)
}
