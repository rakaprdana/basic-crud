package models

type Book struct {
	IDBook          int    `json:"id_book"`
	TitleBook       string `json:"title_book"`
	Author          string `json:"author"`
	DescriptionBook string `json:"description_book"`
	CategoryBook    string `json:"category_book"`
	Cupboard        string `json:"cupboard"`
}
