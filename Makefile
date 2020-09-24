#
all:
	mkdir -p bin
	go build -o bin ./cmd/...

build-lambda:
	mkdir -p bin/lambda
	GOOS=linux go build -o bin/lambda ./lambda/...

.aws-sam:
	sam build --base-dir lambda/md5str --template lambda/md5str/template.yaml

invoke-demo: .aws-sam
	@echo EXPECTED: \"$(shell /bin/echo -n "hello" | md5)\"
	@echo
	echo '{"str": "hello"}' | sam local invoke -e -

clean:
	rm -rf bin
	rm -rf .aws-sam

