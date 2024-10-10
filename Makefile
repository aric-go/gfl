.PHONY: build

build:
	@./scripts/bump_version.sh
	go build -ldflags "-X 'cmd.Version=$(cat VERSION)'" -o dist/gflv main.go
