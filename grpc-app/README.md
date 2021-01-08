## GRPC APP

### How to generate proto
clone submodule googleapis
```sh
git submodule update --init --recursive
```
then go generate
```sh
cd proto/greeter
protoc -I. -I../googleapis --go_out=plugins=grpc:. *.proto
protoc -I. -I../googleapis --grpc-gateway_out=:. *.proto
```
