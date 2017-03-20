package main

import (
    "log"
    "fmt"
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
    serverPort := 7777
    serverPortStr := fmt.Sprintf(":%d", serverPort)
    serverTransport, err := thrift.NewTServerSocket(serverPortStr)
    if err != nil {
        log.Fatal("unable to create server socket ", err)
    }
    processor := buf.NewBufProcessor(new (BufHandler))
    fmt.Println("listening on port=", serverPort)
    server := thrift.NewTSimpleServer2(processor, serverTransport)
    if err = server.Serve(); err != nil {
        log.Fatal(err)
    }
}
