module github.com/ErenDursun/grpc-server-template

go 1.18

replace github.com/grpc-ecosystem/go-grpc-middleware => github.com/ErenDursun/go-grpc-middleware v0.0.0-20221214210458-593df4cfde10

require (
	github.com/golang-jwt/jwt/v4 v4.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	google.golang.org/grpc v1.52.0
)

require (
	cloud.google.com/go/compute/metadata v0.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
