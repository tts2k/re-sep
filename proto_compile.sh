#!/bin/bash
# proto validate should be used with bufbuild but I guess this is fine for now
# Content
rm ./server/content/internal/proto/*
protoc --go_out=./server/content/internal/proto --go_opt=paths=source_relative \
 --go-grpc_out=./server/content/internal/proto --go-grpc_opt=paths=source_relative \
 -I=./proto/ -I=./proto/protovalidate ./proto/*.proto

# User
rm ./server/user/internal/proto/*
protoc --go_out=./server/user/internal/proto --go_opt=paths=source_relative \
 --go-grpc_out=./server/user/internal/proto --go-grpc_opt=paths=source_relative \
 -I=./proto/ -I=./proto/protovalidate/ ./proto/*.proto

# Client
rm -r ./client/src/lib/proto/*
cd ./client && protoc --plugin=$(npm root)/.bin/protoc-gen-ts_proto \
 --ts_proto_out=./src/lib/proto \
 --ts_proto_opt=outputServices=grpc-js \
 --ts_proto_opt=esModuleInterop=true \
 -I=../proto/ -I=../proto/protovalidate/ ../proto/*.proto
