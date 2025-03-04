package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := os.Getenv("connect_db")
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := DB.Ping(); err != nil {
		fmt.Println("Erro de conexão:", err)
		panic(err)
	} else {
		fmt.Println("Conexão bem-sucedida!")
	}
}

func QueryDB(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Função para fazer inserção no banco de dados
func InsertDB(query string, args ...interface{}) (sql.Result, error) {
	result, err := DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
