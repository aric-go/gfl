_default:
    @just --list

# Build ldflags
ldflags := '-s -w'
name := "gfl"

# aliases
alias b := build-unix-like
alias u := linux-upload-server

init-tag:
  git tag v1.0.0 && git push origin v1.0.0

# Run the app (e.g. just start --version)
[group('dev')]
start *args:
    go run main.go {{args}}

# Tidy go modules
[group('dev')]
tidy:
    go mod tidy

# ------- macos start -------
# Build for current platform (macOS)
[group('macos')]
macos-build:
    go build -ldflags "{{ldflags}}" -o dist/{{name}} main.go

# Link the macOS binary into /usr/local/bin
[group('macos')]
macos-link:
    ln -sf $(pwd)/dist/{{name}} /usr/local/bin/{{name}}

# Remove the linked binary from /usr/local/bin
[group('macos')]
macos-unlink:
    rm -f /usr/local/bin/{{name}}

# Upload the macOS binary to Weiyun
[group('macos')]
macos-upload-to-weiyun:
    wy upload dist/{{name}} --dir /100-WeiyunCli/002-bins/000-macos/

# ------- macos end -------

# ------- linux start -------
# Build for Linux amd64
[group('linux')]
linux-build:
    GOOS=linux GOARCH=amd64 go build -ldflags "{{ldflags}}" -o dist/{{name}}-linux-amd64 main.go

# Upload the Linux binary to a server (default: aly1)
[group('linux')]
linux-upload-server server='aly1':
    scp dist/{{name}}-linux-amd64 {{server}}:/usr/local/bin/{{name}}

# Upload the Linux binary to Weiyun
[group('linux')]
linux-upload-to-weiyun:
    wy upload dist/{{name}}-linux-amd64 --dir /100-WeiyunCli/002-bins/001-linux/

# ------- linux end -------


# ------- windows start -------
# Build for Windows amd64
[group('windows')]
windows-build:
    GOOS=windows GOARCH=amd64 go build -ldflags "{{ldflags}}" -o dist/{{name}}-windows-amd64.exe main.go

# Upload the Windows binary to Weiyun
[group('windows')]
windows-upload-to-weiyun:
    wy upload dist/{{name}}-windows-amd64.exe --dir /100-WeiyunCli/002-bins/002-windows/
# ------- windows end -------


# ----- composite start -----
build-unix-like: macos-build linux-build
build-all: build-unix-like windows-build
# ----- composite end -----
