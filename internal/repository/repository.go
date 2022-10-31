package repository

import (
	"github.com/jmoiron/sqlx"
)

type Admin interface {
	Confirmation(transactionId int, value bool) error
}

type Reserve interface {
}

type User interface {
	Transfer(userId int, cash int, userIdTransfer int) error
	GetBalance(userId int) (int, error)
	AddMoney(userId int, cash int) error
	Pay(userId int, cash int, description string) error
}

type Repository struct {
	Admin
	Reserve
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:  NewUserPostgres(db),
		Admin: NewAdminPostgres(db),
	}

}
