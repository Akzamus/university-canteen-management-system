package user

import (
	"context"
	"encoding/json"
	"fmt"
	def "github.com/Akzamus/university-canteen-management-system/api_gateway_service/internal/service"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/pkg/client"
	"github.com/Akzamus/university-canteen-management-system/api_gateway_service/pkg/dto"
	"net/http"
)

var _ def.UserService = (*service)(nil)

type service struct {
	restClient client.RestClient
}

func NewService(userServiceBaseUrl string) *service {
	return &service{
		restClient: client.NewRestClient(userServiceBaseUrl),
	}
}

const (
	getUserByIdPath           = "/api/v1/users/%s"
	getAllUsersPath           = "/api/v1/users/"
	createUserPath            = "/api/v1/users/"
	updateUserPath            = "/api/v1/users/%s"
	deleteUserByIdPath        = "/api/v1/users/%s"
	verifyUserCredentialsPath = "/api/v1/users/verify"
)

func (s *service) GetUserByID(_ context.Context, userUuid string) (dto.UserResponseDto, error) {
	path := fmt.Sprintf(getUserByIdPath, userUuid)

	resp, err := s.restClient.SendRequest(http.MethodGet, path, nil)
	if err != nil {
		return dto.UserResponseDto{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var userResponse dto.UserResponseDto
		err := json.NewDecoder(resp.Body).Decode(&userResponse)
		if err != nil {
			return dto.UserResponseDto{}, fmt.Errorf("failed to decode response body: %w", err)
		}
		return userResponse, nil
	case http.StatusNotFound:
		return dto.UserResponseDto{}, fmt.Errorf("user not found")
	default:
		return dto.UserResponseDto{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func (s *service) GetAllUsers(_ context.Context) ([]dto.UserResponseDto, error) {
	resp, err := s.restClient.SendRequest(http.MethodGet, getAllUsersPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var users []dto.UserResponseDto
		err := json.NewDecoder(resp.Body).Decode(&users)
		if err != nil {
			return nil, fmt.Errorf("failed to decode response body: %w", err)
		}
		return users, nil
	default:
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func (s *service) CreateUser(_ context.Context, userDto dto.UserRequestDto) (dto.UserResponseDto, error) {
	resp, err := s.restClient.SendRequest(http.MethodPost, createUserPath, userDto)
	if err != nil {
		return dto.UserResponseDto{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusCreated:
		var createdUser dto.UserResponseDto
		err := json.NewDecoder(resp.Body).Decode(&createdUser)
		if err != nil {
			return dto.UserResponseDto{}, fmt.Errorf("failed to decode response body: %w", err)
		}
		return createdUser, nil
	default:
		return dto.UserResponseDto{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func (s *service) UpdateUser(_ context.Context, userDto dto.UserRequestDto, uuid string) (dto.UserResponseDto, error) {
	path := fmt.Sprintf(updateUserPath, uuid)

	resp, err := s.restClient.SendRequest(http.MethodPut, path, userDto)
	if err != nil {
		return dto.UserResponseDto{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var updatedUser dto.UserResponseDto
		err := json.NewDecoder(resp.Body).Decode(&updatedUser)
		if err != nil {
			return dto.UserResponseDto{}, fmt.Errorf("failed to decode response body: %w", err)
		}
		return updatedUser, nil
	default:
		return dto.UserResponseDto{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func (s *service) DeleteUserByID(_ context.Context, uuid string) error {
	path := fmt.Sprintf(deleteUserByIdPath, uuid)

	resp, err := s.restClient.SendRequest(http.MethodDelete, path, nil)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusNotFound:
		return fmt.Errorf("user not found")
	default:
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func (s *service) VerifyUserCredentials(_ context.Context, userCredentialsDto dto.UserCredentialsRequestDto) (dto.UserResponseDto, error) {
	resp, err := s.restClient.SendRequest(http.MethodPost, verifyUserCredentialsPath, userCredentialsDto)
	if err != nil {
		return dto.UserResponseDto{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var userResponse dto.UserResponseDto
		err := json.NewDecoder(resp.Body).Decode(&userResponse)
		if err != nil {
			return dto.UserResponseDto{}, fmt.Errorf("failed to decode response body: %w", err)
		}
		return userResponse, nil
	case http.StatusNotFound:
		return dto.UserResponseDto{}, fmt.Errorf("incorrect email or password")
	default:
		return dto.UserResponseDto{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}
