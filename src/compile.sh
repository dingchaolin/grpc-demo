#!/usr/bin/env bash
./compiler/protoc --go_out=. --plugin=protoc-gen-grpc=./compiler/protoc_gen_go ./proto/myproto.proto
./compiler/protoc --grpc_out=. --plugin=protoc-gen-grpc=./compiler/protoc_gen_go  ./proto/myproto.proto

./compiler/protoc --go_out=plugins=grpc:. ./proto/myproto.proto



