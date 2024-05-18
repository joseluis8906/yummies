.PHONY: pb
pb: pbgo pbjs
	@echo "generating pbs ..."

.PHONY: pbgo
pbgo:
	@protoc --proto_path=protobuf/yummies\
		--go_out=go-code/pkg/pb --go_opt=paths=source_relative\
		--go-grpc_out=go-code/pkg/pb --go-grpc_opt=paths=source_relative\
		protobuf/yummies/home.proto


.PHONY: pbjs
pbjs:
	@protoc --proto_path=protobuf/yummies\
		--js_out=import_style=commonjs:js-code/pb\
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:js-code/pb\
		protobuf/yummies/home.proto
