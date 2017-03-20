
package main

import (
  "github.com/sanjosh/golang/thriftproxy/buf/buf"
  "git.apache.org/thrift.git/lib/go/thrift"
  "math/rand"
  "time"
  "fmt"
  "net"
  "log"
)

func main() {

  filename := "testme"
  host := "localhost"
  portStr := "8888"

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

  value := make([]byte, 4096)
  rand.Read(value)

  startTime := time.Now()
  for i := 0; i < 1000; i = i + 1 {
    err := client.WriteData(filename, string(value[:]))
    if err != nil {
        log.Fatal("failed in iteration=%d", i)
    }
  }
  elapsed := time.Since(startTime)
  fmt.Println("time taken ", elapsed)

  data, err := client.ReadData(filename)
  if data != string(value[:]) {
    fmt.Println("error in value")
  }
}
