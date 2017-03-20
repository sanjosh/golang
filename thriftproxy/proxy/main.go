package main

import (
    "log"
    "git.apache.org/thrift.git/lib/go/thrift"
    "github.com/sanjosh/golang/thriftproxy/buf/buf"
    "fmt"
)

var server *thrift.TSimpleServer
var client *buf.BufClient

// Implement buf.buf interface with Write method
type BufHandler struct {
}

type WriteCmd struct {
    filename string 
    data string
}

var writeRespChan chan(error)
var writeCmdChan chan(*WriteCmd)

// this go routine sends writeCmds to backend server
// and reads back results
func backgroundWrite() {
    for {
       writeCmd := <- writeCmdChan
       err := client.WriteData(writeCmd.filename, writeCmd.data)
       writeRespChan <- err 
    }
}

func (*BufHandler) WriteData(filename string, data string) error {
    writeCmdChan <- &WriteCmd{filename, data}
    err := <- writeRespChan
    return err
}

func (*BufHandler) ReadData(filename string) (string, error) {
    data, err := client.ReadData(filename)
    return data, err
}

// Proxy runs on 8888 and talks to server on 7777
func main() {

    writeRespChan = make (chan error)    
    writeCmdChan = make (chan *WriteCmd)    

    backendPort := 7777
    backendPortStr := fmt.Sprintf(":%d", backendPort)

    // start client
    //transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
    protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
    socket, err := thrift.NewTSocket(backendPortStr)
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

    go backgroundWrite()

    // start server
    proxyPort := 8888;
    proxyPortStr := fmt.Sprintf(":%d", proxyPort)
    serverTransport, err := thrift.NewTServerSocket(proxyPortStr)
    if err != nil {
        log.Fatal("unable to create server socket ", err)
    }
    // bind handler
    processor := buf.NewBufProcessor(new (BufHandler))
    fmt.Println("listening on port=", proxyPort);
    server := thrift.NewTSimpleServer2(processor, serverTransport)
    if err = server.Serve(); err != nil {
        log.Fatal(err)
    }
}
