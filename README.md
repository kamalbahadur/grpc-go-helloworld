Protoc
---
bin/compile.sh

Docker Build
---
docker build -t grpc-go-helloworld .

Docker Run
---
docker run -p 50080:50080 -p 51080:51080 grpc-go-helloworld