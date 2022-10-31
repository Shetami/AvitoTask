package repository

import (
	"AvitoTask/internal/models"
	repository "AvitoTask/report"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AdminPostgres struct {
	db *sqlx.DB
}

func NewAdminPostgres(db *sqlx.DB) *AdminPostgres {
	return &AdminPostgres{db: db}
}

func (a *AdminPostgres) Confirmation(transactionId int, value bool) error {
	var transaction models.Buffer
	tr, err := a.db.Begin()
	if err != nil {
		return err
	}
	if value == true {
		q1 := fmt.Sprintf("SELECT balance, user_id, description FROM %s WHERE id=$1", buffer)
		q2 := fmt.Sprintf("DELETE FROM %s WHERE id=$1", buffer)
		q3 := fmt.Sprintf("UPDATE %s SET balance = balance + $1", avito)

		if err := a.db.Get(&transaction, q1, transactionId); err != nil {
			return err
		}

		_, err := tr.Exec(q2, transactionId)
		if err != nil {
			tr.Rollback()
			return err
		}

		_, err = tr.Exec(q3, transaction.Balance)
		if err != nil {
			tr.Rollback()
			return err
		}
		repository.Report(transaction.UserId, transaction.Description, transaction.Balance)
	} else {
		q1 := fmt.Sprintf("UPDATE %s SET balance = balance + $1", users)
		q2 := fmt.Sprintf("DELETE FROM %s WHERE id=$1", buffer)

		_, err = tr.Exec(q1, transaction.Balance)
		if err != nil {
			tr.Rollback()
			return err
		}
		_, err := tr.Exec(q2, transactionId)
		if err != nil {
			tr.Rollback()
			return err
		}
	}
	return tr.Commit()
}
