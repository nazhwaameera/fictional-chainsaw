package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "Bolaitubundar123"
	DB_NAME     = "challenge7"
)

var (
	db  *sql.DB
	err error
)

type Book struct {
	ID          int
	Title       string
	Author      string
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s
						 dbname=%s sslmode=disable`, DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database.")

	// CreateBook()
	// GetAllBooks()
	// UpdateBookbyID()
	DeleteBookbyID()
}

func CreateBook() {
	var book = Book{}

	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	Returning *
	`

	err = db.QueryRow(sqlStatement, "Six Crimson Cranes", "Elizabeth Lim", "A book about cranes").Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Book Data: %+v\n", book)
}

func GetAllBooks() {
	var results = []Book{}

	sqlStatement := `SELECT * from books`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	fmt.Println("Book datas: ", results)
}

func UpdateBookbyID() {
	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1;
	`

	res, err := db.Exec(sqlStatement, 1, "Six of Crows", "Leigh Bardugo", "A book about crows")

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Updataed data amount:", count)
}

func DeleteBookbyID() {
	sqlStatement := `
	DELETE from books
	WHERE id = $1;
	`

	res, err := db.Exec(sqlStatement, 1)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount: ", count)
}
