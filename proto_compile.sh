#!/bin/bash
# Content
protoc --go_out=./server/content/internal/ --go_opt=paths=source_relative \
 --go-grpc_out=./server/content/internal/ --go-grpc_opt=paths=source_relative \
 ./proto/*.proto

# User
protoc --go_out=./server/user/internal/ --go_opt=paths=source_relative \
 --go-grpc_out=./server/user/internal/ --go-grpc_opt=paths=source_relative \
 ./proto/*.proto

# Client
cd ./client && protoc --plugin=$(npm root)/.bin/protoc-gen-ts_proto \
 --ts_proto_out=./src/lib/proto \
 --ts_proto_opt=outputServices=grpc-js \
 --ts_proto_opt=esModuleInterop=true \
 -I=../proto/ ../proto/*.proto
