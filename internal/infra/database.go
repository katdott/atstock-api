package infra

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Importa o driver silenciosamente
)

func ConnectDB() (*sql.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL não encontrada no arquivo .env")
	}

	// Abre a conexão com o PostgreSQL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	// Faz um Ping para garantir que as credenciais estão corretas e a nuvem responde
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}