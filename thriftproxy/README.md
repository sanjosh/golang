go get github.com/sanjosh/golang

# ref to https://github.com/pinterest/bender/blob/master/thrift/TUTORIAL.md
thrift --out src/thriftproxy/buf --gen go:package_prefix=thriftproxy/buf src/thriftproxy/buf/buf.thrift 
go install github.com/sanjosh/golang/thriftproxy/buf/buf
go install github.com/sanjosh/golang/thriftproxy/bufserver
go install github.com/sanjosh/golang/thriftproxy/proxy
go install github.com/sanjosh/golang/thriftproxy/buf/buf/buf-remote

bin/bufserver

bin/proxy

bin/buf-remote -p 8888 WriteData filename "hello world"
