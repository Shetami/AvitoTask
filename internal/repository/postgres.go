package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	users  = "users"
	avito  = "avito"
	buffer = "buffer"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	NameDB   string
	SSLMode  string
}

func NewPSQL(conf Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.NameDB,
		conf.Password,
		conf.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
