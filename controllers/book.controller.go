package controllers

import (
	"fmt"
	"log"

	"github.com/rakaprdana/basic-crud/config"
	"github.com/rakaprdana/basic-crud/models"
)

func AddBook() {
	_, err := config.ConnectDB().Exec("INSERT INTO books(title_book, author, description_book, category_book, cupboard) VALUES ('The Great Gatsby', 'F. Scott Fitzgerald', 'Novel klasik tentang mimpi Amerika', 'Fiction', 'Shelf A')")
	if err != nil {
		log.Fatal("Created error: ", err)
	}

	fmt.Println("Book has been added")
}

func ReadBook() {
	rows, err := config.ConnectDB().Query("SELECT title_book, author, category_book, cupboard FROM books")

	if err != nil {
		log.Fatal("Readed error: ", err)
	}

	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.TitleBook, &book.Author, &book.CategoryBook, &book.Cupboard)

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println("Title: , Author, Category: , Cupboard: \n", book.TitleBook, book.Author, book.CategoryBook, book.Cupboard)
	}
}

func UpdateBook() {
	_, err := config.ConnectDB().Exec("UPDATE books SET author=?, cupboard=? WHERE id_book=?", "Raka", "Shelf B", 2)
	if err != nil {
		log.Fatal("Updated error", err)
	}

	fmt.Println("Updated has been success")
}
func DeletedBook() {
	_, err := config.ConnectDB().Exec("DELETE FROM books WHERE id_book=?", 2)
	if err != nil {
		log.Fatal("Delete error: ", err)
	}

	fmt.Println("Book has been deleted")
}
