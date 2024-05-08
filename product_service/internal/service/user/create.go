package product

import (
	"context"
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/converter"
	"github.com/Akzamus/university-canteen-management-system/product_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) CreateProduct(ctx context.Context, productDto dto.ProductRequestDto) (dto.ProductResponseDto, error) {
	product := converter.ToProductModel(&productDto)

	CreateProduct, err := s.ProductRepository.Save(ctx, product)
	if err != nil {
		return dto.ProductResponseDto{}, errors.Wrap(err, "failed to create product")
	}

	return converter.ToProductResponseDto(&CreateProduct), nil
}
