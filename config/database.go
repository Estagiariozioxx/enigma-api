package config

import (
    "fmt"
    "log"
    "os"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sqlx.DB

func ConnectDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar .env")
    }

    dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
        os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
        os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

    db, err := sqlx.Connect("postgres", dbURL)
    if err != nil {
        log.Fatalf("erro de conexão: %v", err)
    }

    DB = db
    log.Println("Conectado ao database")
}
