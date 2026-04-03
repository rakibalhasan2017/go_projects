package models

import (
	"backend/database"
	"time"
)

// Book struct matches your Postgres table columns
type Book struct {
    ID            int       `db:"id" json:"id"`
    Title         string    `db:"title" json:"title"`
    Author        string    `db:"author" json:"author"`
    PublishedDate *string    `db:"published_date" json:"published_date"`
    Price         float64   `db:"price" json:"price"`
    ImageURL      string    `db:"image_url" json:"image_url"`
    CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

func GetAllBooks() ([]Book, error) {
    books := []Book{}
    err := database.DB.Select(&books, "SELECT * FROM books ORDER BY id DESC")
    return books, err
}

func CreateBook(book *Book) error {
    query := `INSERT INTO books (title, author, published_date, price, image_url) 
              VALUES (:title, :author, :published_date, :price, :image_url) 
              RETURNING id`
    // NamedQuery is great because it maps struct fields to :name markers automatically
    rows, err := database.DB.NamedQuery(query, book)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Next() {
        return rows.Scan(&book.ID)
    }
    return nil
}

func UpdateBook(id int, book *Book) error {
    book.ID = id
    query := `UPDATE books SET title=:title, author=:author, published_date=:published_date, 
              price=:price, image_url=:image_url WHERE id = :id`
    _, err := database.DB.NamedExec(query, book)
    return err
}

func DeleteBook(id int) error {
    _, err := database.DB.Exec("DELETE FROM books WHERE id = $1", id)
    return err
}

// Add Delete and Update methods here following the same pattern...