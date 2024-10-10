.PHONY: build

build:
	@./scripts/bump_version.sh
	go build -ldflags "-X 'github.com/aric-go/github-flow/cmd.Version=$(cat VERSION)'" -o dist/gflv main.go

