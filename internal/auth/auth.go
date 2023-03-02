package auth

import (
	"context"

	grpc_jwt "github.com/ErenDursun/go-grpc-jwt-middleware/jwt"
	"github.com/golang-jwt/jwt/v4"
)

type GrpcServerClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GetClaims(ctx context.Context) *GrpcServerClaims {
	token, ok := ctx.Value(grpc_jwt.DefaultContextKey).(*jwt.Token)
	if !ok {
		return nil
	}

	claims, ok := token.Claims.(*GrpcServerClaims)
	if !ok {
		return nil
	}

	return claims
}
