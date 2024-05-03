package user

import (
	"context"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/converter"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) CreateUser(ctx context.Context, userDto dto.UserRequestDto) (dto.UserResponseDto, error) {
	userModel := converter.ToUserModel(&userDto)

	createdUser, err := s.userRepository.Save(ctx, userModel)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to create user")
	}

	return converter.ToUserResponseDto(&createdUser), nil
}
