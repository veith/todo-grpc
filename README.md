# todo-grpc

## Prereqs

```
brew install protobuf

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go

go get -u moul.io/protoc-gen-gotemplate

go get github.com/gogo/protobuf/protoc-gen-gogofast
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto

```

OR to build protobuf from source, the following tools are needed:

- autoconf
- automake
- libtool
- make
- g++
- unzip

https://github.com/protocolbuffers/protobuf/blob/master/src/README.md


Install protobuf
```
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
```


To use regenerateProto.sh install
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/micro/protobuf/protoc-gen-go

go get github.com/gogo/protobuf/protoc-gen-gogofast 
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto

```
