# Info

this is base thrift demo
- at `thrift version 0.11.0`
- golang version `go1.10.2 darwin/amd64`

## Warning

0.11.0 thrift has same error

- see [https://issues.apache.org/jira/browse/THRIFT-4447?jql=text%20~%20%22thrift.TClient%22](https://issues.apache.org/jira/browse/THRIFT-4447?jql=text%20~%20%22thrift.TClient%22)

# run

```sh
# see help
$ make help

# gen thrift with golang
$ make genThriftGo
```

# Install ENV of thrift

```sh
# macOS
$ brew install thrift
# linux
$ curl -O http://mirrors.tuna.tsinghua.edu.cn/apache/thrift/0.11.0/thrift-0.11.0.tar.gz
$ tar zxvf thrift-0.11.0.tar.gz
$ cd thrift-0.11.0
$ ./bootstrap.sh
$ ./configure
$ make -j 2
$ sudo make install
```

- windows

download http://www.apache.org/dyn/closer.cgi?path=/thrift/0.11.0/thrift-0.11.0.exe
then add PATH

## golang thrift depends

```sh
﻿go get -v git.apache.org/thrift.git/lib/go/thrift/...
```

> this is use thrift golang must have

