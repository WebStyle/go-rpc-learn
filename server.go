package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Person struct {
	Name string
}

type MyClass int

func (t *MyClass) Mul(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *MyClass) SayHello(p *Person, hello *string) error {
	*hello = "Hello " + p.Name
	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error ", err.Error())
	}
}

func main() {
	myclass := new(MyClass)
	rpc.Register(myclass)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		checkError(err)
		rpc.ServeConn(conn)
	}
}
