package repository

import (
	"AvitoTask/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (s *UserPostgres) GetBalance(userId int) (int, error) {
	var balance models.User
	q := fmt.Sprintf("SELECT balance FROM %s WHERE id=$1", users)
	row := s.db.QueryRow(q, userId)
	if err := row.Scan(&balance.Balance); err != nil {
		return 0, err
	}

	return balance.Balance, nil

}

func (s *UserPostgres) Transfer(userId int, cash int, userIdTransfer int) error {
	var balance_first models.User
	var balance_second models.User

	tr, err := s.db.Begin()
	if err != nil {
		return nil
	}

	q := fmt.Sprintf("SELECT balance FROM %s WHERE id=$1", users)
	row1 := tr.QueryRow(q, userId)
	if err := row1.Scan(&balance_first.Balance); err != nil {
		tr.Rollback()
		return nil
	}

	row2 := tr.QueryRow(q, userIdTransfer)

	if err := row2.Scan(&balance_second.Balance); err != nil {
		tr.Rollback()
		return nil
	}

	result := balance_first.Balance - cash
	if result <= 0 {
		tr.Rollback()
		return nil
	}
	add_result := balance_second.Balance + cash

	q2 := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id = $2", users)
	_, err = tr.Exec(q2, result, userId)
	if err != nil {
		tr.Rollback()

		return nil
	}

	_, err = tr.Exec(q2, add_result, userIdTransfer)
	if err != nil {
		tr.Rollback()
		return nil
	}

	return tr.Commit()
}

func (s *UserPostgres) AddMoney(userId int, cash int) error {

	q := fmt.Sprintf("UPDATE %s SET balance=balance + $1 WHERE id = $2", users)
	_, err := s.db.Exec(q, cash, userId)
	return err
}

func (s *UserPostgres) Pay(userId int, cash int, description string) error {
	var balance models.User
	var remainder int

	tr, err := s.db.Begin()
	if err != nil {
		return nil
	}

	q := fmt.Sprintf("SELECT balance FROM %s WHERE id=$1", users)
	row1 := tr.QueryRow(q, userId)
	if err := row1.Scan(&balance.Balance); err != nil {
		tr.Rollback()
		return nil
	}

	remainder = balance.Balance - cash
	if remainder < 0 {
		tr.Rollback()
		return nil
	}

	q2 := fmt.Sprintf("INSERT INTO %s (balance, description, user_id) VALUES($1, $2, $3)", buffer)
	_, err = tr.Exec(q2, cash, description, userId)
	if err != nil {
		tr.Rollback()
		return nil
	}

	q3 := fmt.Sprintf("UPDATE %s SET balance=$1 WHERE id=$2", users)
	_, err = tr.Exec(q3, remainder, userId)
	if err != nil {
		tr.Rollback()
		return nil
	}

	return tr.Commit()
}
