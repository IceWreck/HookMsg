run:
	go run ./cmd/hookmsg

run-sample:
	go run ./cmd/sample

build:
	go build -o ./bin/HookMsg -tags hooktelegram ./cmd/hookmsg
