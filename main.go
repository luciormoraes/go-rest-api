package main

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {

	models.Celebrities = []models.Celebrity{
		{Name: "Name 1", History: "History 1"},
		{Name: "Name 2", History: "History 2"},
	}

	fmt.Println("Starting server...")

	routes.HandleRequest()
}
