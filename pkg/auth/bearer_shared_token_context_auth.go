package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"shorturl/pkg/params"
	"strings"
)

type BearerSharedTokenContextAuth struct {
	token string
}

func NewBearerSharedTokenContextAuth(params params.Params) BearerSharedTokenContextAuth {
	token := params.Get("SHARED_TOKEN")
	if token == "" {
		panic("shared token is not found in parameters list")
	}
	return BearerSharedTokenContextAuth{
		token: token,
	}
}

func (b BearerSharedTokenContextAuth) Authorize(context *gin.Context) error {
	authHeader := context.Request.Header.Get("authorization")
	if authHeader == "" {
		return errors.New("no authorization header is provided")
	}
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return errors.New("invalid authorization header format - bearer is expected")
	}
	if authHeaderParts[1] != b.token {
		return errors.New("invalid token")
	}
	return nil
}
