run:
	go run ./cmd/hookmsg

build:
	CGO_ENABLED=0 go build \
		-o ./bin/HookMsgComplete ./cmd/hookmsg

# required when you need to build with older glibc (for older servers)
start-old-container:
	podman run -it --rm -v $(shell pwd):/src:z docker.io/library/golang:1.14.15-stretch