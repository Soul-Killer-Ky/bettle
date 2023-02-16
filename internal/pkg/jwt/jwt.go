package jwt

import (
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

const reason string = "AUTH_FAILURE"

var (
	ErrSignToken = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
)

// Option is jwt option.
type Option func(*options)

// Parser is a jwt parser
type options struct {
	signingMethod jwtv4.SigningMethod
	claims        func() jwtv4.Claims
}

func WithSigningMethod(method jwtv4.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

func WithClaims(f func() jwtv4.Claims) Option {
	return func(o *options) {
		o.claims = f
	}
}

type CustomUserClaims struct {
	ID uint `json:"id"`
	jwtv4.RegisteredClaims
}

func WithCustomUserClaims(uid uint, exp int64) Option {
	now := time.Now()
	claims := CustomUserClaims{
		uid,
		jwtv4.RegisteredClaims{
			IssuedAt:  jwtv4.NewNumericDate(now),
			NotBefore: jwtv4.NewNumericDate(now),
			ExpiresAt: jwtv4.NewNumericDate(now.Add(time.Duration(exp) * time.Second)),
		},
	}
	return WithClaims(func() jwtv4.Claims {
		return claims
	})
}

func CreateToken(key string, opts ...Option) (string, error) {
	claims := jwtv4.RegisteredClaims{}
	o := &options{
		signingMethod: jwtv4.SigningMethodHS256,
		claims:        func() jwtv4.Claims { return claims },
	}
	for _, opt := range opts {
		opt(o)
	}
	token := jwtv4.NewWithClaims(o.signingMethod, o.claims())
	tokenStr, err := token.SignedString([]byte(key))
	if err != nil {
		return "", ErrSignToken
	}

	return tokenStr, nil
}
