
# https://github.com/pinterest/bender/blob/master/thrift/TUTORIAL.md
thrift --out src/thriftproxy/buf --gen go:package_prefix=thriftproxy/buf src/thriftproxy/buf/buf.thrift 
