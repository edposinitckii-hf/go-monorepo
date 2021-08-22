-include ../../Makefile.go.mk

.PHONY: test-unit test-integration build

build:
	CGO := 0
	DIR_OUT := $(CURDIR)/out
	BINARY := echo-api
	BINARY_SRC := $(CURDIR)/cmd/api/...
	call(build-image)
