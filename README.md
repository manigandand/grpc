# gRPC

https://grpc.io/
## Requirements

```
pipenv shell
brew install protobuf
OR
https://github.com/protocolbuffers/protobuf/blob/master/src/README.md
https://askubuntu.com/questions/1072683/how-can-i-install-protoc-on-ubuntu-16-04
sudo apt-get update
sudo apt-get install protobuf-compiler

# Python
pip install grpcio grpcio-tools

# Go
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

# Node
npm install -g grpc-tools

# If permission issue
sudo chown -R $(whoami) ~/.npm
sudo chown -R $(whoami) /usr/lib/node_modules
sudo chown -R $(whoami) /usr/local/lib/node_modules

```
