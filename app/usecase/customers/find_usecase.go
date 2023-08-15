package customers

import (
	"context"
	"github.com/oktopriima/dimy/app/models"
)

func (c *customerUseCase) FindUseCase(ctx context.Context, ID int) (*models.Customer, error) {
	output, err := c.repo.Find(ctx, ID)
	if err != nil {
		return nil, err
	}

	return output, nil
}
