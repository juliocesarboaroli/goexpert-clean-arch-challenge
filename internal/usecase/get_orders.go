package usecase

import "github.com/juliocesarboaroli/goexpert-clean-arch-challenge/internal/entity"

type (
	OrderOutput struct {
		ID         string  `json:"id"`
		Price      float64 `json:"price"`
		Tax        float64 `json:"tax"`
		FinalPrice float64 `json:"final_price"`
	}

	GetOrdersUseCase struct {
		OrderRepository entity.OrderRepositoryInterface
	}
)

func NewGetOrdersUseCase(repository entity.OrderRepositoryInterface) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: repository,
	}
}

func (g *GetOrdersUseCase) Execute() []OrderOutput {
	orders := g.OrderRepository.FindAll()

	return mapResult(orders)
}

func mapResult(orders []entity.Order) (result []OrderOutput) {
	for _, o := range orders {
		result = append(result, OrderOutput{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}

	return result
}
