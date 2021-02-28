# Build Types
# a) Hooks only
# b) Hooks and Matrix
# c) Hooks, Matrix and Telegram

run:
	go run -tags exclude_tg ./cmd/hookmsg

run-sample:
	go run ./cmd/sample

run-newfilter:
	go run ./cmd/newfilter

build:
	go build -o ./bin/HookMsgComplete ./cmd/hookmsg


# required when you need to build with older glibc (for older servers)
start-old-container:
	podman run -it --rm -v $(shell pwd):/src:z docker.io/library/golang:1.14.15-stretch