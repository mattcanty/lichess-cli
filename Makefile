test:
	go test
build: test
	go build
install: build
	go install
