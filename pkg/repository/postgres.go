package repository

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "content.users"
	todoListsTable  = "content.todo_lists"
	usersListsTable = "content.users_lists"
	todoItemsTable  = "content.todo_items"
	listsItemsTable = "content.lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := sqlx.Open("postgres", uri)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//migration, err := migrate.New("file://migrations", uri)
	//if err != nil {
	//	return nil, err
	//}
	//if err := migration.Up(); err != nil {
	//	log.Info().Msgf("migrations: %s", err.Error())
	//}

	return db, nil
}
