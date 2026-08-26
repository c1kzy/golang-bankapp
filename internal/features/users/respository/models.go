package core_http_respository

import "github.com/c1kzy/golang-bankapp/internal/core/domain"

type UserModel struct {
	ID      int
	Version int

	FullName string
	Balance  int
}

func userModelToDomain(user UserModel) domain.User {
	return domain.NewUser(
		user.ID,
		user.Version,
		user.FullName,
		user.Balance,
	)

}

func userModelsToDomain(users []UserModel) ([]domain.User, error) {
	usersDomain := make([]domain.User, len(users))

	for i, user := range users {
		usersDomain[i] = userModelToDomain(user)
	}

	return usersDomain, nil
}
