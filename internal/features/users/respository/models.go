package core_http_respository

import "time"

type UserModel struct {
	ID      int
	Version int

	FullName string
	Balance  int

	CreatedAt time.Time
}
