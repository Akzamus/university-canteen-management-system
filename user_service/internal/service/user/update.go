package user

import (
	"context"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/converter"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/utils/crypto/bcrypt"
	"github.com/Akzamus/university-canteen-management-system/user_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) UpdateUser(ctx context.Context, userDto dto.UserRequestDto, uuid string) (dto.UserResponseDto, error) {
	user := converter.ToUserModel(&userDto)
	user.Uuid = uuid

	hashedPassword, err := bcrypt.GenerateHash(userDto.Password)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to hash password")
	}
	user.Password = hashedPassword

	updatedUser, err := s.userRepository.Save(ctx, user)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to update user")
	}

	return converter.ToUserResponseDto(&updatedUser), nil
}
