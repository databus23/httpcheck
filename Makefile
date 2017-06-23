
VERSION  := 0.1.0
PACKAGES := $(shell find . -type d)
GOFILES  := $(addsuffix /*.go,$(PACKAGES))
GOFILES  := $(wildcard $(GOFILES))

GOFLAGS=-ldflags="-s -w"


build: bin/darwin/httpcheck

bin/%: $(GOFILES) Makefile
	GOOS=$(*D) GOARCH=amd64 go build $(GOFLAGS) -v -i -o $(@D)/$(@F) .

.PHONY: dist
dist: bin/linux/httpcheck bin/darwin/httpcheck bin/windows/httpcheck.exe
	mkdir -p  release/
	rm -rf release/*
	cp bin/darwin/httpcheck release/httpcheck_darwin
	cp bin/linux/httpcheck release/httpcheck_linux
	cp bin/windows/httpcheck.exe release/httpcheck.exe

.PHONY: release
release: dist
ifndef GITHUB_ACCESS_TOKEN
	$(error GITHUB_ACCESS_TOKEN is undefined)
endif
	git push
	gh-release create databus23/httpcheck $(VERSION) master
