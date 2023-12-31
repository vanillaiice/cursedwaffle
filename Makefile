run:
	go run .

build:
	go build -ldflags="-s -w" .

# Compile for Windows
windows:
	mkdir -p bin/windows
	GOOS=windows go build -ldflags="-s -w" -o bin/windows/cursedwaffle.exe cmd/cursedwaffle/main.go

# Compile for Linux
linux:
	mkdir -p bin/linux
	GOOS=linux go build -ldflags="-s -w" -o bin/linux/cursedwaffle cmd/cursedwaffle/main.go

all: linux windows
