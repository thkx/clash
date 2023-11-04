package auth

import (
	"github.com/hhhhhhhxx/clash/component/auth"
)

var authenticator auth.Authenticator

func Authenticator() auth.Authenticator {
	return authenticator
}

func WithAuthenticator(au auth.Authenticator) {
	authenticator = au
}
