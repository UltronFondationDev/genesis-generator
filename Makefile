.PHONY: build

build:
	GOPROXY=$(GOPROXY) \
	go build -o build/genesis-generator ./main.go