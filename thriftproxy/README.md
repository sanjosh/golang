# Download
* go get github.com/sanjosh/golang

# Build
* thrift --out src/thriftproxy/buf --gen go:package_prefix=thriftproxy/buf src/thriftproxy/buf/buf.thrift 
* go install github.com/sanjosh/golang/thriftproxy/buf/buf
* go install github.com/sanjosh/golang/thriftproxy/bufserver
* go install github.com/sanjosh/golang/thriftproxy/proxy
* go install github.com/sanjosh/golang/thriftproxy/client

# Run
* bin/bufserver
* bin/proxy
* bin/client -p 8888 -n 1000 -s 4096

[golang thrift reference](https://github.com/pinterest/bender/blob/master/thrift/TUTORIAL.md)
