.PHONY: help

help:
	@./scripts/make-help.pl

clean:
	@echo "-------------------"
	@echo "delete generated artifacts"
	@echo "-------------------"
	@go clean
	@rm -fr target || true

grpc: query_grpc.pb.go query.pb.go
	@echo "-------------------"
	@echo "generate protobuf service files"
	@echo "-------------------"
	@protoc --proto_path=proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative query.proto

docker:
	@echo "-------------------"
	@echo "build docker image"
	@echo "-------------------"
	@docker build --file Dockerfile --tag grahamcrowell/query:latest .
	@docker push grahamcrowell/query:latest

format:
	@echo "-------------------"
	@echo "format source code"
	@echo "-------------------"
	@gofmt -s -w -l -e **/*.go
	@gofmt -s -w -l -e *.go
	@go vet

pre-commit:
	@echo "-------------------"
	@echo "pre-commit validation"
	@echo "-------------------"
	@# https://stackoverflow.com/a/67962664/5154695
	@# test -z $(gofmt -l **/*.go)
	@./scripts/fmt-err.sh
	@make test

test:
	@echo "-------------------"
	@echo "execute unit and integration tests"
	@echo "-------------------"
	@go test *.go;
