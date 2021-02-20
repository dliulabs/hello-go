# Git setup

```
git init
git add .
git commit -m "update repo"
git branch -M main
git remote rm origin
git remote add origin git@github.com:dliulabs/hello-go.git
git push -u origin main
```

# Run

```
export PATH=$PATH:$(go env GOPATH)/bin
go get
go run main.go hello.go
```

# Build

```
go build
./hello
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