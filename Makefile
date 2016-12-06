.PHONY: all
all:
	@echo "hi"

.PHONY: bindir
bindir:
	@mkdir -p ./bin

.PHONY: build
build: bindir
	GOGC=off go build -i -gcflags="-e" -o ./bin/ehpi ./cmd/ehpi

.PHONY: run
run:
	fresh -p ./cmd/ehpi -r '-config=./etc/ehpi.conf' -o ./bin/ehpi

.PHONY: tools
tools:
	go get -u github.com/pressly/fresh
	go get -u github.com/kardianos/govendor
