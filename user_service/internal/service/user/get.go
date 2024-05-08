package user

import (
	"context"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/converter"
	"github.com/Akzamus/university-canteen-management-system/user_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) GetUserByID(ctx context.Context, uuid string) (dto.UserResponseDto, error) {
	user, err := s.userRepository.FindById(ctx, uuid)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to get user by ID")
	}
	return converter.ToUserResponseDto(&user), nil
}
