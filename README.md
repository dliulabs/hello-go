# Run

```
export PATH=$PATH:$(go env GOPATH)/bin
go get
go run main.go hello.go
```

# Build

```
go build
```

# Install

```
go install
bash -c '$(go env GOPATH)/bin/hello'
```

# Test

```
go test 
go test -v
```

# Colorize Tests

```
export PATH=$PATH:$(go env GOPATH)/bin
go get -u github.com/rakyll/gotest
gotest -v
```