package product

import (
	"context"
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/converter"
	"github.com/Akzamus/university-canteen-management-system/product_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) GetAllProduct(ctx context.Context) ([]dto.ProductResponseDto, error) {
	products, err := s.ProductRepository.FindAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all products")
	}
	return converter.ToProductResponseDtoList(&products), nil
}
