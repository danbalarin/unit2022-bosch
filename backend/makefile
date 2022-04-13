cmdWebserver := ./cmd/web

export ENV=development

all: dev

build:
	go build ${cmdWebserver}

dev: wire
	nodemon --exec 'go run ${cmdWebserver} || exit 1' --watch ./ --ext go --signal SIGTERM

deploy: test build
	# DEPLOY

test:
	go test ./...

wire:
	wire ${cmdWebserver}/wire.go

install:
	rm ${cmdWebserver}/wire_gen.go || true
	go mod tidy
	go install github.com/google/wire/cmd/wire@latest
	make wire