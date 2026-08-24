package domain

type User struct {
	ID      int
	Version int

	FullName string
	Balance  int
}

func NewUnitializedUser(fullName string, balance int) User {
	return User{
		ID:       UnitializedID,
		Version:  UnitializedVersion,
		FullName: fullName,
		Balance:  balance,
	}
}

func NewUser(
	id int,
	version int,
	fullName string,
	balance int,
) User {
	return User{
		ID:       id,
		Version:  version,
		FullName: fullName,
		Balance:  balance,
	}
}
