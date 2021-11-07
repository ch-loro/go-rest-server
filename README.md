# go-server

## Launch server

Launch the following commands:

```bash
go run cmd/sample-server/main.go --port=8000
```

## Request

Access to the following URI:

```bash
http://127.0.0.1:8000/v1/PATH
```

## NOTE

### Generate server codes

```bash
swagger validate swagger.yaml
swagger generate server -f swagger.yaml
```

### Initialize go.mod and install packages

```bash
go mod init
go mod tidy
go build ./...
```
