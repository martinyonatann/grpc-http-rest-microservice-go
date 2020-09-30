protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:pkg/api todo-service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/api todo-service.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --swagger_out=logtostderr=true:api/swagger todo-service.proto