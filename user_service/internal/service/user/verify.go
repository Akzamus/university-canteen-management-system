package user

import (
	"context"
	"fmt"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/converter"
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/utils/crypto/bcrypt"
	"github.com/Akzamus/university-canteen-management-system/user_service/pkg/dto"
	"github.com/pkg/errors"
)

func (s *service) VerifyUserCredentials(ctx context.Context, userCredentialsDto dto.UserCredentialsRequestDto) (dto.UserResponseDto, error) {
	user, err := s.userRepository.FindByEmail(ctx, userCredentialsDto.Email)
	if err != nil {
		return dto.UserResponseDto{}, errors.Wrap(err, "failed to get user by email")
	}

	if !bcrypt.CompareHashAndString(user.Password, userCredentialsDto.Password) {
		return dto.UserResponseDto{}, fmt.Errorf("incorrect email or password")
	}

	return converter.ToUserResponseDto(&user), nil
}
