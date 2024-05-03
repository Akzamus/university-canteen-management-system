package user

import (
	"context"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/converter"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) UpdateUser(ctx context.Context, userDto dto.UserRequestDto) (dto.UserResponseDto, error) {
	user := converter.ToUserModel(&userDto)

	updatedUser, err := s.userRepository.Save(ctx, user)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to update user")
	}

	return converter.ToUserResponseDto(&updatedUser), nil
}
