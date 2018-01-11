# protoc --proto_path=. --java_out=java --python_out=python --go_out=go grpc.proto
python -m grpc_tools.protoc -I grpc --python_out=grpc --grpc_python_out=grpc grpc.proto
protoc -I grpc grpc/grpc.proto --go_out=plugins=grpc:grpc
# protoc --proto_path=. --java_out=java --grpc_java_out=java grpc.proto
