# Build Linux & Windows binaries
export GOOS=linux
go run build/build.go
export GOOS=windows
go build -o bin/chrise.exe e.go