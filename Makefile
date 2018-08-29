.PHONY: pkg

build:
	mkdir -p bin
	GOOS=darwin go build -o bin/aws-osx cmd/aws/main.go
	GOOS=darwin go build -o bin/local-osx cmd/local/main.go
	GOOS=linux go build -o bin/aws-linux cmd/aws/main.go
	GOOS=linux go build -o bin/local-linux cmd/local/main.go

pkg:
	cd bin && zip raffleChecker.zip aws-osx local-osx aws-linux local-linux

clean:
	rm -rf bin/*

release: clean build pkg