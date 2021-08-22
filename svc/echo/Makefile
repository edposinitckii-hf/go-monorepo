-include ../../Makefile.go.mk

.PHONY: test-unit test-integration build

build: CGO = 0
build: DIR_OUT=$(CURDIR)/out
build:
	$(call build-image,echo-api,$(CURDIR)/cmd/api/...)
