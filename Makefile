gen: 
	protoc --proto_path=./protobuf --go_out=./generate --go-grpc_out=./generate ./protobuf/admin.proto
	protoc --proto_path=./protobuf --go_out=./generate --go-grpc_out=./generate ./protobuf/customer.proto
	protoc --proto_path=./protobuf --go_out=./generate --go-grpc_out=./generate ./protobuf/delivery.proto
	protoc --proto_path=./protobuf --go_out=./generate --go-grpc_out=./generate ./protobuf/notification.proto
