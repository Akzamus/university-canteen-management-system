package product

import (
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/repository"
	def "github.com/Akzamus/university-canteen-management-system/product_service/internal/service"
)

var _ def.ProductService = (*service)(nil)

type service struct {
	ProductRepository repository.ProductRepository
}

func NewService(ProductRepository repository.ProductRepository) *service {
	return &service{
		ProductRepository: ProductRepository,
	}
}
