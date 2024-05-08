package converter

import (
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/model"
	"github.com/Akzamus/university-canteen-management-system/product_service/pkg/dto"
)

func ToProductResponseDto(product *model.Product) dto.ProductResponseDto {
	return dto.ProductResponseDto{
		Uuid:  product.Uuid,
		Name: product.Name,
		Price: product.Price, 
	}
}

func ToProductModel(productDto *dto.ProductRequestDto) model.Product {
	return model.Product{
		Name: productDto.Name,
		Description: productDto.Description,
	}
}

func ToProductResponseDtoList(products *[]model.Prodct) []dto.ProductResponseDto {
	ProductDtoList := make([]dto.ProductResponseDto, len(*products))
	for i, product := range *&products {
		ProductDtoList[i] = ToProductResponseDto(&product)
	}
	return ProductDtoList
}
