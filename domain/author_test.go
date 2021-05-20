package domain_test

import (
	"blog-golang/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	author := domain.Author{
		Name:     "Author",
		Email:    "author@mail.com",
		Password: "_authorpass",
		Avatar:   "https://avatar.png",
	}

	err := author.Create()
	require.Nil(t, err)

	author.Avatar = "invalid_url"
	err = author.Create()
	require.Error(t, err)
	author.Avatar = "https://avatar.png"

	author.Email = "invalidmail.com"
	err = author.Create()
	require.Error(t, err)
	author.Email = "author@mail.com"

	author.Name = ""
	fmt.Println(author)
	err = author.Create()
	require.Error(t, err)
	author.Name = "Author Name"

	author.Password = ""
	err = author.Create()
	require.Error(t, err)
}
