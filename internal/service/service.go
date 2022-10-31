package service

import (
	"AvitoTask/internal/repository"
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

type Service struct {
	Admin
	Reserve
	User
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		User:  NewUserService(rep.User),
		Admin: NewAdminService(rep.Admin),
	}

}
