package user

import (
	"context"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/converter"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/internal/utils/crypto/bcrypt"
	"github.com/Akzam/usuniversity-canteen-management-system/user_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) CreateUser(ctx context.Context, userDto dto.UserRequestDto) (dto.UserResponseDto, error) {
	user := converter.ToUserModel(&userDto)

	hashedPassword, err := bcrypt.GenerateHash(userDto.Password)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to hash password")
	}
	user.Password = hashedPassword

	createdUser, err := s.userRepository.Save(ctx, user)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to create user")
	}

	return converter.ToUserResponseDto(&createdUser), nil
}
