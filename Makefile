echo:
	echo "Hello"

build:
	go build -o bin/pdfreader src/main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/pdfreader-arm src/main.go
	GOOS=linux GOARCH=arm64 go build -o bin/pdfreader-arm64 src/main.go
	GOOS=linux GOARCH=386 go build -o bin/pdfreader-386 src/main.go

all: echo build
