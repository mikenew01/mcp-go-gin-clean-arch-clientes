package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func NewPostgresDB() *sql.DB {
	connStr := "user=mikemcp password=M@ikonCan123 dbname=clean-arch sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o banco de dados PostgresDB com sucesso!")

	return db
}
