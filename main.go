package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {

	// Mock personalities
	models.Personalities = []models.Personality{
		{ID: 1, Name: "Name 1", History: "History 1"},
		{ID: 2, Name: "Name 2", History: "History 2"},
		{ID: 3, Name: "Name 3", History: "History 3"},
	}

	database.ConnectDB()

	fmt.Println("Starting server...")

	routes.HandleRequest()

}
