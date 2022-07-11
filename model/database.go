package model

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(127.0.0.1:3306)/db_library?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db
}

func CreateDatabase() {
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_library")
	err := db.Ping()
	db.Close()
	if err != nil {
		db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
		createTables(db)
		db.Close()
		fmt.Println("Database created sucessfuly!")
	} else {
		fmt.Println("Database already exists!")
	}
}

func createTables(db *sql.DB) {
	execute(db, "CREATE DATABASE IF NOT EXISTS db_library;")

	execute(db, "USE db_library;")

	execute(db, `

		CREATE TABLE books
		(
			id 				INTEGER,
			name 			VARCHAR(80),
			description 	VARCHAR(100),

			CONSTRAINT PK_tb_books_id PRIMARY KEY (id)
		);`,
	)

	execute(db, `CREATE TABLE categories
		(
			id 				INTEGER,
			name 			VARCHAR(100),

			CONSTRAINT PK_tb_categories_id PRIMARY KEY (id)
		);`,
	)
	execute(db, `CREATE TABLE bookshelves
		(
			book 				INTEGER,
			category 			INTEGER,

			CONSTRAINT PK_tb_bookshelf_book_category PRIMARY KEY (book, category),
			CONSTRAINT FK_tb_bookshelf_tb_books_id FOREIGN KEY (book)
				REFERENCES books (id),
			CONSTRAINT FK_tb_bookshelf_tb_category_id FOREIGN KEY (category)
				REFERENCES categories (id)
		);`,
	)
}
func execute(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}
