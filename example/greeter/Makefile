.PHONY: build

build:
	protoc \
	  	-I. \
		-I../../include/graphql \
		--plugin=../../dist/protoc-gen-graphql \
		--go-grpc_out=./greeter \
		--graphql_out=verbose:./greeter \
		greeter.proto
