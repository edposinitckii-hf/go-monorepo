vendor:
	go mod vendor

test-unit: vendor
	go test -v -tags unit -short -race ./...

test-integration: vendor
	go test -v -run Integration -tags integration -race ./...

define build-image
	CGO_ENABLED=${CGO} go build -o ${DIR_OUT}/$1 ${GO_LINKER_FLAGS} $2
endef