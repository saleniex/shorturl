package auth

import "github.com/gin-gonic/gin"

// ContextAuth interface provides authorization from GIN context
type ContextAuth interface {
	// Authorize performs authorization using credentials provided either in headers or in request data
	// In case authorization fails method returns error object
	Authorize(context *gin.Context) error
}
