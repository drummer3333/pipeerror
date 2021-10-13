#!/bin/bash
export BUILD_VERSION="ERROR-TEST-1.0"

echo "run go get on main.go"
(cd cmd/test/ ; go get)

echo "install goimports"
go install golang.org/x/tools/cmd/goimports@v0.1.7

echo "manually get joncalhoun/pipe (imports of generator are ignored)"
go get github.com/joncalhoun/pipe

echo "run generator"
go generate ./internal/version/helper.go

echo "run main.go"
(cd cmd/test/ ; go run main.go)