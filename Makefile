# Build Featuresets:
# a) Scripts only
# b) Scripts and Matrix
# c) Scripts, Matrix and Telegram

# feature set b
run:
	go run -tags matrix,script ./cmd/hookmsg

run-sample:
	go run -tags matrix,script ./cmd/sample

run-newfilter:
	go run -tags matrix,script ./cmd/newfilter

# feature set c
build:
	go build \
		-tags matrix,telegram,script \
		-o ./bin/HookMsgComplete ./cmd/hookmsg

# feature set a
build-a:
	go build \
		-tags script \
		-o ./bin/HookMsg_A ./cmd/hookmsg

# feature set b
build-b:
	go build \
		-tags script,matrix \
		-o ./bin/HookMsg_A ./cmd/hookmsg


# required when you need to build with older glibc (for older servers)
start-old-container:
	podman run -it --rm -v $(shell pwd):/src:z docker.io/library/golang:1.14.15-stretch