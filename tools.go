//go:build tools
// +build tools

package main

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate go install github.com/golang/protobuf/protoc-gen-go

//go:generate protoc -I proto --go_opt=paths=source_relative --go_out=proto/src/go --go-grpc_opt=paths=source_relative --go-grpc_out=proto/src/go proto/otp_service.proto
//go:generate protoc -I proto --go_out=proto/src/go --go_opt=paths=source_relative --go-grpc_out=proto/src/go --go-grpc_opt=paths=source_relative proto/auth_service.proto

import (
	_ "github.com/favadi/protoc-go-inject-tag"
	_ "github.com/protoc-gen-micro"
	_ "github.com/protoc/v3"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
