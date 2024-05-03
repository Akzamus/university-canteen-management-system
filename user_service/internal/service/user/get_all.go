package user

import (
	"context"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/converter"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) GetAllUsers(ctx context.Context) ([]dto.UserResponseDto, error) {
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all users")
	}
	return converter.ToUserResponseDtoList(&users), nil
}
