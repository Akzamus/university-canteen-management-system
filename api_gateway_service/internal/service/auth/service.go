package auth

import (
	"context"
	"fmt"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/converter"
	def "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/pkg/dto"
	"github.com/go-chi/jwtauth"
)

var _ def.AuthService = (*service)(nil)

type service struct {
	jwtAuth     *jwtauth.JWTAuth
	userService def.UserService
}

func NewService(jwtAuth *jwtauth.JWTAuth, userService def.UserService) *service {
	return &service{
		jwtAuth:     jwtAuth,
		userService: userService,
	}
}

func (s *service) Register(ctx context.Context, authDto dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	user := converter.ToUserRequestDto(&authDto)

	createdUser, err := s.userService.CreateUser(ctx, user)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("failed to register user: %w", err)
	}

	claims := map[string]interface{}{
		"userUuid": createdUser.Uuid,
		"role":     createdUser.Role,
	}

	_, tokenString, _ := s.jwtAuth.Encode(claims)

	return dto.AuthResponseDto{AccessToken: tokenString}, nil
}

func (s *service) Authenticate(ctx context.Context, authDto dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	userCredentials := converter.ToUserCredentialsRequestDto(&authDto)

	user, err := s.userService.VerifyUserCredentials(ctx, userCredentials)
	if err != nil {
		return dto.AuthResponseDto{}, fmt.Errorf("failed to authenticate user: %w", err)
	}

	claims := map[string]interface{}{
		"userUuid": user.Uuid,
		"role":     user.Role,
	}

	_, tokenString, _ := s.jwtAuth.Encode(claims)

	return dto.AuthResponseDto{AccessToken: tokenString}, nil
}
