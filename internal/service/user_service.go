package service

import (
	"AvitoTask/internal/repository"
)

type UserService struct {
	rep repository.User
}

func NewUserService(rep repository.User) *UserService {
	return &UserService{rep: rep}
}

func (s *UserService) GetBalance(userId int) (int, error) {
	return s.rep.GetBalance(userId)
}

func (s *UserService) Transfer(userId int, cash int, userIdTransfer int) error {
	return s.rep.Transfer(userId, cash, userIdTransfer)
}

func (s *UserService) AddMoney(userId int, cash int) error {
	return s.rep.AddMoney(userId, cash)
}

func (s *UserService) Pay(userId int, cash int, description string) error {
	return s.rep.Pay(userId, cash, description)
}
