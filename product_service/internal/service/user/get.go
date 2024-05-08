package product

import (
	"context"

	"github.com/Akzamus/university-canteen-management-system/product_service/internal/converter"
	"github.com/Akzamus/university-canteen-management-system/product_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) GetProductByID(ctx context.Context, uuid string) (dto.ProductrResponseDto, error) {
	user, err := s.ProductRepository.FindById(ctx, uuid)
	if err != nil {
		return dto.ProductResponseDto{}, errors.Wrap(err, "failed to get product by ID")
	}
	return converter.ToProdcutResponseDto(&user), nil
}
