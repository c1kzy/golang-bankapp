package core_user_transport_grpc

import (
	"fmt"

	"github.com/c1kzy/golang-bankapp/generated/users"
	"github.com/c1kzy/golang-bankapp/internal/core/domain"
)

func newUserFromMessage(req *users.CreateUserRequest) (domain.User, error) {
	if &req.Balance == nil {
		return domain.User{}, fmt.Errorf("balance cannot be nil")
	}

	newUser := domain.NewUnitializedUser(
		req.FullName,
		int(req.Balance),
	)

	return newUser, nil
}

func DTOtoUserPatch(req *users.PatchUserRequest) domain.UserPatch {
	balance := int(req.Balance)
	return domain.UserPatch{
		FullName: req.FullName,
		Balance:  &balance,
	}
}
