package main

import (
	"strings"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"log"
	"context"
	"github.com/sinlov/thrift-golang-demo/example"
	"time"
)

type FormatDataImpl struct {
}

func (fdi *FormatDataImpl) DoFormat(ctx context.Context, data *example.Data) (r *example.Data, err error) {
	var rData example.Data
	rData.Text = strings.ToUpper(data.Text)
	fmt.Printf("[ %v ] FormatDataImpl.DoFormat ctx => %v, data => %v, return %v\n", time.Now().String(), ctx, data, rData)
	return &rData, nil
}

const (
	HOST = "localhost"
	PORT = "28080"
)

func main() {

	handlerIml := &FormatDataImpl{
	}
	processor := example.NewFormatDataProcessor(handlerIml)

	// init thrift server
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST+":"+PORT)
	server.Serve()
}
