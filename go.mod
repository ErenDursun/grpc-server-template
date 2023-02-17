module github.com/ErenDursun/grpc-server-template

go 1.18

replace github.com/grpc-ecosystem/go-grpc-middleware => github.com/ErenDursun/go-grpc-middleware v0.0.0-20221214210458-593df4cfde10

require (
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	google.golang.org/grpc v1.53.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
