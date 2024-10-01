package usecase

import (
	"github.com/ankardo/CleanArch/internal/dto"
	"github.com/ankardo/CleanArch/internal/entity"
)

type FindAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindAllOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *FindAllOrdersUseCase {
	return &FindAllOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *FindAllOrdersUseCase) Execute() ([]dto.OrderOutputDTO, error) {
	outputData, err := c.OrderRepository.FindAll()
	if err != nil {
		return []dto.OrderOutputDTO{}, err
	}
	var output []dto.OrderOutputDTO
	for _, data := range outputData {
		order := dto.OrderOutputDTO{
			ID:         data.ID,
			Price:      data.Price,
			Tax:        data.Tax,
			FinalPrice: data.FinalPrice,
		}
		output = append(output, order)
	}
	return output, nil
}
