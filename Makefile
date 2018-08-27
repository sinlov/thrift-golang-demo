.PHONY: dist test build

genThriftGo:
	cd ./thrift_file &&  sh gen-thrift-go.sh

runExampleService:
	cd ./exampleserver && go run server.go

runExampleClient:
	cd ./exampleclient && go run client.go

help:
	@echo "make genThriftGo -> gen thrift file at folder thrift_file"
	@echo "make runExampleService -> run thrift example server"
	@echo "make runExampleClient -> run thrift example client"