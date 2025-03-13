package erro

import (
	"errors"
)

var (
	ErrStatusUnauthorized 	= errors.New("invalid Token")
	ErrTokenExpired		 	= errors.New("token expired")
	ErrBadRequest		 	= errors.New("internal error")
	ErrNotFound 			= errors.New("data not found")
	ErrCertRevoked			= errors.New("error cert revoke")
	ErrParseCert			= errors.New("error parse cert")
)