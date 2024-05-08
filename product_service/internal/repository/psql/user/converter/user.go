package converter

import (
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/model"
	"github.com/Akzamus/university-canteen-management-system/product_service/internal/repository/psql/user/entity"
)

func ToProductModel(product *entity.Product) model.Product {
	return model.Product{
		Uuid:     product.Uuid,
		Name:     product.Name,
		Description: product.Description,
		price:     product.Price,
		Image_url:	product.Image_url,
		Category_id: product.Category_id,
		Total_quantity: product.Total_quantity,
	}
}

func ToProductModelList(products *[]entity.User) []model.User {
	servProduct := make([]model.User, len(*&products))
	for i, product := range *&products {
		servProduct[i] = ToProductModel(&product)
	}
	return servProduct
}

func ToProductEntity(product *model.Product) entity.Product {
	return entity.Product{
		Uuid:     product.Uuid,
		Name:     product.Name,
		Description: product.Description,
		price:     product.Price,
		Image_url:	product.Image_url,
		Category_id: product.Category_id,
		Total_quantity: product.Total_quantity,
	}
}
