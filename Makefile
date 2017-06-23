
PACKAGES := $(shell find . -type d)
GOFILES  := $(addsuffix /*.go,$(PACKAGES))
GOFILES  := $(wildcard $(GOFILES))

build: bin/darwin/httpcheck

bin/%: $(GOFILES) Makefile
	GOOS=$(*D) GOARCH=amd64 go build $(GOFLAGS) -v -i -o $(@D)/$(@F) .
