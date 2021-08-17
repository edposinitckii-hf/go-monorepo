vendor:
	go mod vendor

test-unit: vendor
	go test -v -tags unit -short -race ./...

test-integration: vendor
	go test -v -run Integration -tags integration -race ./...