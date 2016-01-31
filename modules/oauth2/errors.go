package oauth2

import (
	"github.com/Gr1N/pacman/modules/errors"
)

var (
	errServiceNotSupported = errors.New(
		"authorization_service_not_supported", "Specified service is not supported")
	errServiceUnavailable = errors.New(
		"authorization_service_unavailable", "Authorization service temporary unavailable")
	errStateInvalid = errors.New(
		"state_invalid", "State value is not valid")
	errCodeInvalid = errors.New(
		"authorization_code_invalid", "Authorization code invalid or issued for another service")
	errTokenInvalid = errors.New(
		"access_token_invalid", "Access token invalid or expired")
)
