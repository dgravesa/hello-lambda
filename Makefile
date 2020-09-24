#
all:
	mkdir -p bin
	go build -o bin ./cmd/...

build-lambda:
	mkdir -p bin-lambda
	GOOS=linux go build -o bin-lambda ./lambda/...

clean:
	rm -rf bin
	rm -rf bin-lambda

