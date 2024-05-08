package user

import (
	"github.com/Akzamus/university-canteen-management-system/user_service/internal/repository"
	def "github.com/Akzamus/university-canteen-management-system/user_service/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}
