package models

type Book struct {
	IDBook   int    `json:"id_book"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Cupboard string `json:"cupboard"`
}
