package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	maxOpenConns    = 25
	maxIdleConns    = 5
	connMaxLifetime = 5 * time.Minute
	connMaxIdleTime = 10 * time.Minute
)

// PostgresDB обертка над *sql.DB
type PostgresDB struct {
	DB *sql.DB
}

// NewPostgresDB создает новое подключение к PostgreSQL
func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	// Проверка подключения
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logrus.Info("Successfully connected to PostgreSQL database")

	return &PostgresDB{DB: db}, nil
}

// Close закрывает соединение с базой данных
func (p *PostgresDB) Close() error {
	if p.DB != nil {
		logrus.Info("Closing database connection")
		return p.DB.Close()
	}
	return nil
}

// GetDB возвращает *sql.DB для прямого использования
func (p *PostgresDB) GetDB() *sql.DB {
	return p.DB
}

