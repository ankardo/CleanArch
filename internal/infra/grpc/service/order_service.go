package service

import (
	"context"

	"github.com/ankardo/CleanArch/internal/dto"
	"github.com/ankardo/CleanArch/internal/infra/grpc/pb"
	"github.com/ankardo/CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase   usecase.CreateOrderUseCase
	FindAllOrdersUseCase usecase.FindAllOrdersUseCase
	FindOrderUseCase     usecase.FindOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	findAllOrdersUseCase usecase.FindAllOrdersUseCase,
	findOrderUseCase usecase.FindOrderUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase:   createOrderUseCase,
		FindAllOrdersUseCase: findAllOrdersUseCase,
		FindOrderUseCase:     findOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := dto.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) FindAllOrders(ctx context.Context, in *pb.Blank) (*pb.FindAllOrdersResponse, error) {
	output, err := s.FindAllOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*pb.FindOrderResponse
	for _, data := range output {
		orders = append(orders, &pb.FindOrderResponse{
			Id:         data.ID,
			Price:      float32(data.Price),
			Tax:        float32(data.Tax),
			FinalPrice: float32(data.FinalPrice),
		})
	}
	return &pb.FindAllOrdersResponse{Orders: orders}, nil
}

func (s *OrderService) FindOrder(ctx context.Context, in *pb.FindOrderRequest) (*pb.FindOrderResponse, error) {
	output, err := s.FindOrderUseCase.Execute(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.FindOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
