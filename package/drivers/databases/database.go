package databases

import (
	"database/sql"
	"fmt"
)

type DatabaseConnection struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
}

type DatabaseRepositoryInterface interface {
	Close() error
}

type DatabaseRepository struct {
	Conn *sql.DB
}

func (db *DatabaseRepository) Close() error {
	return db.Conn.Close()
}

func NewDatabaseRepository(param DatabaseConnection) (DatabaseRepositoryInterface, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", param.Username, param.Password, param.Host, param.Port, param.DatabaseName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, nil
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка при проверке подключения: %w", err)
	}

	return &DatabaseRepository{Conn: db}, nil
}
