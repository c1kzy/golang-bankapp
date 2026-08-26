package core_http_server

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
)

func GetIDPathValue(r *http.Request, name string) (int, error) {
	idPathValue := r.PathValue(name)

	intValue, err := strconv.Atoi(idPathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"invalid id: %w",
			core_errors.ErrInvalidArgument)
	}

	return intValue, nil
}
