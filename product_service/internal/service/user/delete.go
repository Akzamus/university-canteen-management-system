package product

import (
	"context"

	"github.com/pkg/errors"
)

func (s *service) DeleteProductByID(ctx context.Context, uuid string) error {
	err := s.ProductRepository.DeleteById(ctx, uuid)
	if err != nil {
		return errors.Wrap(err, "failed to delete product by ID")
	}
	return nil
}
