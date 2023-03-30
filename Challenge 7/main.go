package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "Bolaitubundar123"
	DB_NAME     = "challenge7-1"
)

var (
	db  *sql.DB
	err error
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []Book `json:"data"`
	Message string `json:"message"`
}

func setupDB() *sql.DB {
	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s
						 dbname=%s sslmode=disable`, DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = sql.Open("postgres", psqlInfo)

	// if err != nil {
	// 	panic(err)
	// }

	// defer db.Close()

	// err = db.Ping()

	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Successfully connected to database.")

	return db
}
func main() {

	router := mux.NewRouter()

	// CreateBook()
	// GetAllBooks()
	// GetBookbyId()
	// UpdateBookbyID()
	// DeleteBookbyID()

	router.HandleFunc("/books", GetAllBooks).Methods("GET")

	router.HandleFunc("/books/{bookID}", GetBookbyId).Methods("GET")

	router.HandleFunc("/books", CreateBook).Methods("POST")

	router.HandleFunc("/books/{bookID}", UpdateBookbyID).Methods("PUT")

	router.HandleFunc("/books/{bookID}", DeleteBookbyID).Methods("DELETE")

	fmt.Println("Server starts at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// fungsi untuk memasukkan data buku baru
func CreateBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")
	description := r.FormValue("description")

	db := setupDB()

	var book = Book{}

	sqlStatement := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	Returning *
	`

	err = db.QueryRow(sqlStatement, title, author, description).Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		panic(err)
	}

	var response = JsonResponse{Type: "success", Message: "The book data has been successfully added."}

	json.NewEncoder(w).Encode(response)
	// fmt.Printf("New Book Data: %+v\n", book)
}

// fungsi untuk mengambil data semua buku
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

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

	var response = JsonResponse{Type: "success", Data: results}

	json.NewEncoder(w).Encode(response)
	// fmt.Println("Book datas: ", results)
}

// fungsi untuk mengambil data semua buku
func GetBookbyId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID := params["bookID"]

	db := setupDB()

	var results = []Book{}

	sqlStatement := `SELECT * from books WHERE id = $1`

	rows, err := db.Query(sqlStatement, bookID)

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
	var response = JsonResponse{Type: "success", Data: results}

	json.NewEncoder(w).Encode(response)

	// fmt.Println("Book datas: ", results)
}

// fungsi untuk mengupdate data buku
func UpdateBookbyID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID := params["bookID"]

	title := r.FormValue("title")
	author := r.FormValue("author")
	description := r.FormValue("description")

	db := setupDB()

	sqlStatement := `
	UPDATE books
	SET title = $2, author = $3, description = $4
	WHERE id = $1;
	`

	_, err := db.Exec(sqlStatement, bookID, title, author, description)

	if err != nil {
		panic(err)
	}

	var response = JsonResponse{Type: "success", Message: "The data has been successfully updated."}

	json.NewEncoder(w).Encode(response)

}

// fungsi untuk menghapus data buku
func DeleteBookbyID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookID := params["bookID"]

	sqlStatement := `
	DELETE from books
	WHERE id = $1;
	`

	_, err := db.Exec(sqlStatement, bookID)

	if err != nil {
		panic(err)
	}

	var response = JsonResponse{Type: "success", Message: "The data has been successfully deleted."}

	json.NewEncoder(w).Encode(response)
}
