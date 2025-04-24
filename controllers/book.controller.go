package controllers

import (
	"fmt"
	"log"

	"github.com/rakaprdana/basic-crud/config"
	"github.com/rakaprdana/basic-crud/models"
)

func AddBook() {
	query := "INSERT INTO books(title, author, category, cupboard) VALUES (?, ?, ?, ?)"
	_, err := config.ConnectDB().Exec(query,
		"The Great Gatsby", "F. Scott Fitzgerald", "Fiction", "Shelf A")
	if err != nil {
		log.Fatal("Created error: ", err)
	}

	fmt.Println("Book has been added")
}

func ReadBook() {
	query := "SELECT title, author, category, cupboard FROM books"
	rows, err := config.ConnectDB().Query(query)

	if err != nil {
		log.Fatal("Readed error: ", err)
	}

	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Title, &book.Author, &book.Category, &book.Cupboard)
		if err != nil {
			log.Println("Row scan error:", err)
			continue
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Fatal("Rows error:", err)
	}

	if len(books) == 0 {
		fmt.Println("No books found.")
		return
	}

	fmt.Println("Books found:")
	for _, book := range books {
		fmt.Println("Title: Author:  Category:  Cupboard: \n", book.Title, book.Author, book.Category, book.Cupboard)
	}
}

func UpdateBook() {
	query := "UPDATE books SET author=?, cupboard=? WHERE id_book=?"
	_, err := config.ConnectDB().Exec(query, "Raka", "Shelf B", 2)
	if err != nil {
		log.Fatal("Updated error", err)
	}

	fmt.Println("Updated has been success")
}

func DeletedBook() {
	query := "DELETE FROM books WHERE id_book=?"
	_, err := config.ConnectDB().Exec(query, 4)
	if err != nil {
		log.Fatal("Delete error: ", err)
	}

	fmt.Println("Book has been deleted")
}

// soft delete methode (optional)
func SoftDeletedBook() {
	query := "UPDATE books SET is_deleted=true WHERE id_book=?"
	_, err := config.ConnectDB().Exec(query, 1)
	if err != nil {
		log.Fatal("Soft delete error: ", err)
	}
	fmt.Println("Book has been deleted by soft delete")
}
