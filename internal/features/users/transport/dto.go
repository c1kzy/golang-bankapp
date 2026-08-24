package core_http_transport

import "github.com/c1kzy/golang-bankapp/internal/core/domain"

func newUserFromDTO(request CreateUserRequest) domain.User {
	return domain.NewUnitializedUser(
		request.FullName,
		request.Balance,
	)
}

func userDomainToDTO(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:        user.ID,
		Version:   user.Version,
		FullName:  user.FullName,
		Balance:   user.Balance,
	}
}
