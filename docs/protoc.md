## How to install `protoc`

- Download latest `protoc` release from [here](https://github.com/protocolbuffers/protobuf/releases)
- We'll have something like `protoc-3.11.4-linux-x86_64.zip` with the following structure:
```text
    bin
        protoc
    include
        google
            protobuf
                ... many files here ...
```
- Place `bin` content into `$PATH`-searchable - `bin`
- Place `include` near `bin`, so we'll have something like the following:
```text
    bin
        ... $PATH-searchable bin folder ...
        ... you may have your old bin files ...
        protoc
    include
        google
            protobuf
                ... many files here ...
``` 

## How to install `protoc-gen-go`
### From sources
The simplest way is to run 
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
The compiler plugin, protoc-gen-go, will be installed in $GOPATH/bin unless $GOBIN is set. It must be in your $PATH for the protocol compiler, protoc, to find it.
### From package
```
sudo apt install golang-goprotobuf-dev
```
After that
```
which protoc-gen-go
/usr/bin/protoc-gen-go
```
