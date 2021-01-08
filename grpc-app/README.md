## GRPC APP

### How to generate proto
```sh
cd proto/greeter
protoc -I. --go_out=plugins=grpc:. *.proto
```
