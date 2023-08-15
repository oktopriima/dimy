package customers

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
	"github.com/oktopriima/dimy/app/repository"
)

type customerUseCase struct {
	repo repository.CustomerRepository
}

type CustomerUseCase interface {
	FindUseCase(ctx context.Context, ID int) (*models.Customer, error)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
