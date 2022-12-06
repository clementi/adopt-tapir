NAME=adopt-tapir

build:
	go build -o bin/${NAME}

clean:
	-rm -rf bin
	
run:
	go run main.go

arch:
	GOOS=darwin  GOARCH=amd64 go build -o bin/darwin-amd64-${NAME}
	GOOS=darwin  GOARCH=arm64 go build -o bin/darwin-arm64-${NAME}
	GOOS=freebsd GOARCH=amd64 go build -o bin/freebsd-amd64-${NAME}
	GOOS=linux   GOARCH=amd64 go build -o bin/linux-amd64-${NAME}
	GOOS=netbsd  GOARCH=amd64 go build -o bin/netbsd-amd64-${NAME}
	GOOS=openbsd GOARCH=amd64 go build -o bin/openbsd-amd64-${NAME}
	GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64-${NAME}.exe
