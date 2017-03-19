package main

import (
    "log"
    "git.apache.org/thrift.git/lib/go/thrift"
    "github.com/sanjosh/golang/thriftproxy/buf/buf"
    _ "fmt"
)

var server *thrift.TSimpleServer
var client *buf.BufClient

// Implement buf.buf interface with Write method
type BufHandler struct {
}

var writeRespChan chan(error)

func (*BufHandler) WriteData(filename string, data string) error {
    //go func () {
        //err := client.WriteData(filename, data)
        //writeRespChan <- err
    //}()
    //err := <- writeRespChan
    //return err
    return client.WriteData(filename, data)
}

func (*BufHandler) ReadData(filename string) (string, error) {
    data, err := client.ReadData(filename)
    return data, err
}

// Proxy runs on 8888 and talks to server on 7777
func main() {

    writeRespChan = make (chan error)    

    // start client
    //transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    socket, err := thrift.NewTSocket(":7777")
    if err != nil {
        log.Fatal("client failed ", err)
    }
    //transport := transportFactory.GetTransport(socket)
    defer socket.Close()
    if err := socket.Open(); err != nil {
        log.Fatal("open failed ", err)
    }

    client = buf.NewBufClientFactory(socket, protocolFactory)

    err = client.WriteData("tmp", "hello")
    if err != nil {
        log.Fatal("test write failed ", err)
    }

    // start server
    serverTransport, err := thrift.NewTServerSocket(":8888")
    if err != nil {
        log.Fatal("unable to create server socket ", err)
    }
    // bind handler
    processor := buf.NewBufProcessor(new (BufHandler))
    server := thrift.NewTSimpleServer2(processor, serverTransport)
    if err = server.Serve(); err != nil {
        log.Fatal(err)
    }
}
