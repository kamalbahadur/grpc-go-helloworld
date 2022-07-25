#!/bin/bash
set -ex

rm -rf proto/helloworld
mkdir proto/helloworld
protoc -I proto --go_out proto --go-grpc_out proto --grpc-gateway_out proto/helloworld --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative proto/*.proto