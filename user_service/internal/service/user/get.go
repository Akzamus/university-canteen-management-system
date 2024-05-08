package user

import (
	"context"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/converter"
	"github.com/Akzamus/university-canteen-management-system/user_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) GetAllUsers(ctx context.Context) ([]dto.UserResponseDto, error) {
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all users")
	}
	return converter.ToUserResponseDtoList(&users), nil
}

func (s *service) GetUserByID(ctx context.Context, uuid string) (dto.UserResponseDto, error) {
	user, err := s.userRepository.FindById(ctx, uuid)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to get user by ID")
	}
	return converter.ToUserResponseDto(&user), nil
}
