# go-grpc-calculator

## Build

```bash
sudo dnf install protobuf
```
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```
```bash
go run server/main.go
```
```bash
go run client/main.go -server="localhost:8080"
```

### gprcui
```bash
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
grpcui --plaintext 127.0.0.1:8080
```

## Resources

- Video: [YouTube](https://www.youtube.com/watch?v=_4TPM6clQjM) [GitHub](https://github.com/dreamsofcode-io/grpc)
