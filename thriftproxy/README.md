
go install github.com/sanjosh/golang/thriftproxy/buf/buf
go install github.com/sanjosh/golang/thriftproxy/bufserver
go install github.com/sanjosh/golang/thriftproxy/proxy
go install github.com/sanjosh/golang/thriftproxy/buf/buf/buf-remote

bin/bufserver

bin/proxy

bin/buf-remote -p 8888 WriteData filename "hello world"
