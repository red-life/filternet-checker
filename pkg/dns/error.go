package dns

import "errors"

var (
	InvalidDomainErr     = errors.New("invalid domain address")
	InvalidServerErr     = errors.New("invalid domain address")
	InvalidServerPortErr = errors.New("invalid server port")
)
