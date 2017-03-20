
package main

import (
  "github.com/sanjosh/golang/thriftproxy/buf/buf"
  "git.apache.org/thrift.git/lib/go/thrift"
  "math/rand"
  "time"
  "fmt"
  "net"
  "log"
  "flag"
  "os"
)

func Usage() {
    fmt.Fprintln(os.Stderr, "usage of ", os.Args[0], "[-h host:port] [-n num_iterations] [-s buffer_size]")
}

func main() {

  flag.Usage = Usage

  var host string
  var port int
  var bufferSize int
  var numIter int

  flag.StringVar(&host, "h", "localhost", "host and port")
  flag.IntVar(&port, "p", 7777, "port")
  flag.IntVar(&numIter, "n", 1000, "num_iterations")
  flag.IntVar(&bufferSize, "s", 4096, "buffer size")
  flag.Parse()
  portStr := fmt.Sprint(port)

  filePrefix := "testme"

  socket, err := thrift.NewTSocket(net.JoinHostPort(host, portStr))
  if err != nil {
    log.Fatal("cannot connect %s %s", host, portStr)
  }
  defer socket.Close()
  protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

  client := buf.NewBufClientFactory(socket, protocolFactory)

  if err := socket.Open(); err != nil {
    log.Fatal("error opening", err)
  }

  value := make([]byte, bufferSize)
  rand.Read(value)

  startTime := time.Now()
  for i := 0; i < numIter; i = i + 1 {
    filename := fmt.Sprintf("%s_%d", filePrefix, i)
    err := client.WriteData(filename, string(value[:]))
    if err != nil {
        log.Fatal("failed in iteration=%d", i)
    }
  }
  elapsed := time.Since(startTime)
  fmt.Println(bufferSize, ",", numIter, ",", elapsed)

  //data, err := client.ReadData(filename)
  //if data != string(value[:]) {
    //fmt.Println("error in value")
  //}
}
