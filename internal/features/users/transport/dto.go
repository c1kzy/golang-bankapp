package core_user_transport

import (
	"fmt"

	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

type UserRequest struct {
	FullName string `json:"full_name"`
	Balance  *int   `json:"balance"  `
}

type UserResponse struct {
	ID       int    `json:"id"`
	Version  int    `json:"version"`
	FullName string `json:"full_name"`
	Balance  *int   `json:"balance"`
}

func newUserFromDTO(request CreateUserRequest) (domain.User, error) {
	if request.Balance == nil {
		return domain.User{}, fmt.Errorf("balance is required")
	}
	newUser := domain.NewUnitializedUser(
		request.FullName,
		*request.Balance,
	)

	return newUser, nil
}

func userDomainToDTO(user domain.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Version:  user.Version,
		FullName: user.FullName,
		Balance:  user.Balance,
	}
}

func usersDomainToDTO(users []domain.User) []UserResponse {
	usersResponse := make([]UserResponse, len(users))
	for i, user := range users {
		usersResponse[i] = userDomainToDTO(user)
	}

	return usersResponse
}

func userDTOtoUserPatch(request UserRequest) domain.UserPatch {
	return domain.UserPatch{
		FullName: request.FullName,
		Balance:  request.Balance,
	}
}
