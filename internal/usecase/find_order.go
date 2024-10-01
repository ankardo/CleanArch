package usecase

import (
	"github.com/ankardo/CleanArch/internal/dto"
	"github.com/ankardo/CleanArch/internal/entity"
)

type FindOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *FindOrderUseCase {
	return &FindOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *FindOrderUseCase) Execute(input string) (dto.OrderOutputDTO, error) {
	outputData, err := c.OrderRepository.Find(input)
	if err != nil {
		return dto.OrderOutputDTO{}, err
	}
	output := dto.OrderOutputDTO{
		ID:         outputData.ID,
		Price:      outputData.Price,
		Tax:        outputData.Tax,
		FinalPrice: outputData.FinalPrice,
	}
	return output, nil
}
