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
cd ../main && go run main.go hello.go
```

# Build

```
go build -o hello
cd ../main && go build -o hello
./hello
```

# Install

```
go get
go install
bash -c '$(go env GOPATH)/bin/hello'
```

# Test

```
go test 
go test -v
go test -run Hello
```

# Colorize Tests

```
export PATH=$PATH:$(go env GOPATH)/bin
go get -u github.com/rakyll/gotest
gotest -v
```

# Test Coverage

```
go test -cover
go test --coverprofile=cover.txt
go tool cover -html=cover.txt -o cover.html
```

# Go Module

```
mkdir main
mv hello* main* main/
mkdir -p nummanip/calc
mv math* nummanip/calc/
```

## Init a Module

```
cd nummanip
go mod init github.com/dliulabs/calc
go tidy
```

## Run Test

```
cd nummanip
go test ./calc -v
```