# getport

A simple Go module that provides methods for find an open port on the local
system. A comprehensive set of methods for getting open TCP or UDP ports on
either the IPv4 or IPv6 network stacks are provided.

## Install

```shell
$ go get github.com/jsumners/go-getport
```

## Example

The following example shows a simple TCP server that will write `hello world`
and close the connection when a connection is made, e.g. via [netcat][netcat]
(`nc 127.0.0.1 <discovered-port>`).

```go
package main

import (
	getport "github.com/jsumners/go-getport"
	"net"
)

func main() {
	// Get an open TCP port on the localhost address.
	// `getport.GetTcpPort()` would instead listen on all addresses with
	// a common free port.
	portResult, getPortError := getport.GetTcpPortForAddress("127.0.0.1")
	if getPortError != nil {
		panic(getPortError)
	}

	listenAddress := getport.PortResultToAddress(portResult)
	listener, listenError := net.Listen("tcp", listenAddress)
	if listenError != nil {
		panic(listenError)
	}

	println("listening on: ", listenAddress)
	defer listener.Close()

	connection, connError := listener.Accept()
	if connError != nil {
		panic(connError)
	}

	connection.Write([]byte("hello world"))
	connection.Close()
}
```

[netcat]: https://en.wikipedia.org/wiki/Netcat
