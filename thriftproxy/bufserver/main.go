package main

import (
    "log"
    _ "fmt"
    "io/ioutil"
    "git.apache.org/thrift.git/lib/go/thrift"
    "github.com/sanjosh/golang/thriftproxy/buf/buf"
)

var server *thrift.TSimpleServer

// Implement buf.buf interface with Write method
type BufHandler struct {
}

func (*BufHandler) WriteData(filename string, data string) error {
    err := ioutil.WriteFile(filename, []byte(data), 0644)
    return err
}

func (*BufHandler) ReadData(filename string) (string, error) {
    data, err := ioutil.ReadFile(filename)
    return string(data[:]), err
}

func main() {
    serverTransport, err := thrift.NewTServerSocket(":7777")
    if err != nil {
        log.Fatal("unable to create server socket ", err)
    }
    processor := buf.NewBufProcessor(new (BufHandler))
    server := thrift.NewTSimpleServer2(processor, serverTransport)
    if err = server.Serve(); err != nil {
        log.Fatal(err)
    }
}
