package service

import "AvitoTask/internal/repository"

type AdminService struct {
	rep repository.Admin
}

func NewAdminService(rep repository.Admin) *AdminService {
	return &AdminService{rep: rep}
}

func (a *AdminService) Confirmation(transactionId int, value bool) error {
	return a.rep.Confirmation(transactionId, value)
}
