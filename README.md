Protoc
---
protoc -I helloworld helloworld/*.proto --go_out=plugins=grpc:helloworld

Docker Build
---
docker build -t grpc-go-helloworld .

Docker Run
---
docker run -p 50080:50080 grpc-go-helloworld