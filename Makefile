.PHONY: gen

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=paths=source_relative:services/common/genproto/orders \
		--go-grpc_out=paths=source_relative:services/common/genproto/orders