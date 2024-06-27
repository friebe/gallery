## Start application

```go
go run cmd/webgallery/main.go
```

Server started at port :8080

## Compile for raspberry pi w zero 

```
GOARCH=arm GOARM=6 GOOS=linux go build -o bin/webgallery ./cmd/webgallery/main.go
```

## Execute webgallery 
```
./bin/webgallery
```