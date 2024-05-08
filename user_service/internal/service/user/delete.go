package user

import (
	"context"
	"github.com/pkg/errors"
)

func (s *service) DeleteUserByID(ctx context.Context, uuid string) error {
	err := s.userRepository.DeleteById(ctx, uuid)
	if err != nil {
		return errors.Wrap(err, "failed to delete user by ID")
	}
	return nil
}
