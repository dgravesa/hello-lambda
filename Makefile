#
all:
	mkdir -p bin
	go build -o bin ./cmd/...

clean:
	rm -r bin

