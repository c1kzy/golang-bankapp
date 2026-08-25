package core_http_server

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/c1kzy/golang-bankapp/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var (
	requestValidator = validator.New(validator.WithRequiredStructEnabled())
)

func DecodeAndValidate(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(&dest); err != nil {
		return fmt.Errorf(
			"unable to decode and validate request %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	if err := requestValidator.Struct(dest); err != nil {
		return fmt.Errorf(
			"validation failed: %v: %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
