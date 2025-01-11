package main

import (
	"github.com/amentobing/todo-app-go/database"
	"github.com/amentobing/todo-app-go/routes"
)

func main() {
	database.Connect()
	routes.Setup()
}