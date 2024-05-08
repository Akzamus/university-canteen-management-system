package product

import (
	"context"

	"github.com/Akzamus/university-canteen-management-system/product_service/internal/converter"
	"github.com/Akzamus/university-canteen-management-system/product_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) UpdateProduct(ctx context.Context, productDto dto.ProductRequestDto, uuid string) (dto.ProductResponseDto, error) {
	product := converter.ToProductModel(&productDto)
	product.Uuid = uuid

	updatedProduct, err := s.userRepository.Save(ctx, product)
	if err != nil {
		return dto.ProductResponseDto{}, errors.Wrap(err, "failed to update product")
	}

	return converter.ToUserResponseDto(&updatedProduct), nil
}
