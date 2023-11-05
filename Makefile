run:
	go run ./cmd/hookmsg

build:
	CGO_ENABLED=0 go build \
		-o ./bin/HookMsgComplete ./cmd/hookmsg
