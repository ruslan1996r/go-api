.PHONY: run
# run
run:
	go build -o ./bin/server ./cmd/server && ./bin/server

.PHONY: run-hot
# run in hot-reload mode
run-hot:
	@air --build.cmd "go build -o ./bin/server ./cmd/server" --build.bin "./bin/server"

.PHONY: docs
# generate swagger docs
docs:
	swag init -g ./cmd/server/main.go -o docs
