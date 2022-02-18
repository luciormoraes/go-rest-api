package main

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{ID: 1, Name: "Name 1", History: "History 1"},
		{ID: 2, Name: "Name 2", History: "History 2"},
	}

	fmt.Println("Starting server...")

	routes.HandleRequest()
}
