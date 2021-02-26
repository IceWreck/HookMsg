run:
	go run -tags exclude_tg ./cmd/hookmsg

run-sample:
	go run ./cmd/sample

build:
	go build -o ./bin/HookMsg ./cmd/hookmsg
