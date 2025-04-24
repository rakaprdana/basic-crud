package main

import (
	"fmt"
	"os"

	"github.com/rakaprdana/basic-crud/config"
	"github.com/rakaprdana/basic-crud/controllers"
)

func main() {
	config.ConnectDB()
	fmt.Print("Database has been connected\n")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [create|read|update|delete|softdelete]")
		return
	}

	switch os.Args[1] {
	case "create":
		controllers.AddBook()
	case "read":
		controllers.ReadBook()
	case "update":
		controllers.UpdateBook()
	case "softdelete":
		controllers.SoftDeletedBook()
	case "delete":
		controllers.DeletedBook()
	}
}
