package main

import (
	"fmt"

	"github.com/delba/api-paris/models"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	err = FetchAndSaveCategories()
	handle(err)

	var categories models.Categories

	err = categories.All()
	handle(err)

	for _, category := range categories {
		fmt.Println(category.ID, category.Name)
	}
}

func FetchAndSaveCategories() error {
	var categories models.Categories

	err := categories.Fetch()
	if err != nil {
		return err
	}

	for _, category := range categories {
		err = category.Save()
		if err != nil {
			return err
		}
	}

	return err
}
