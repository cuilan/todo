BINARY="todo"
VERSION=1.0.0
BUILD=`date +%FT%T%z`

PACKAGES=`go list ./... | grep -v /`
GOFILES=`find . -name "*.go" -type f -not -path ""`

PACKAGES=``

default:
	@go build -o ${BINARY}

run: default
	./${BINARY}