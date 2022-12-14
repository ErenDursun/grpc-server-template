package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth/jwt"
)

type GrpcServerClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GetClaims(ctx context.Context) *GrpcServerClaims {
	token, ok := ctx.Value(grpc_auth.DefaultJWTConfig.ContextKey).(*jwt.Token)
	if !ok {
		return nil
	}

	claims, ok := token.Claims.(*GrpcServerClaims)
	if !ok {
		return nil
	}

	return claims
}
