package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"fmt"
	"log"
	"github.com/sinlov/thrift-golang-demo/example"
	"context"
)

const (
	HOST = "localhost"
	PORT = "28080"
)

var defaultCtx = context.Background()

type clientImpl struct {
	c thrift.TClient
}

func (c *clientImpl) Call(ctx context.Context, method string, args, result thrift.TStruct) error {
	fmt.Printf("clientImpl.Call ctx => %v,  method %v, args => %v, result => %v\n", ctx, method, args, result)
	//protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	return nil
}

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, err := transportFactory.GetTransport(tSocket)
	if err != nil {
		log.Fatalln("transport error:", err)
	}

	//vClientImpl := &clientImpl{
	//
	//}
	//client := example.NewFormatDataClient(vClientImpl)

	// old way, need update?
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := example.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()

	data := example.Data{
		Text: "hello,world!",
	}
	d, err := client.DoFormat(defaultCtx, &data)
	if err != nil {
		log.Fatalln("client DoFormat Error:", HOST+":"+PORT, err)
	}
	fmt.Printf("data => %v\n", d)
}
