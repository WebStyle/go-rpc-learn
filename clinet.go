package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

type Args struct {
	A, B int
}

type Person struct {
	Name string
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error ", err.Error())
	}
}

func main() {

	client, err := rpc.Dial("tcp", "localhost:1234")
	checkError(err)

	a := os.Args[1]
	b := os.Args[2]

	aInt, err := strconv.Atoi(a)
	bInt, err := strconv.Atoi(b)
	checkError(err)

	args := Args{aInt, bInt}
	var reply int

	// Call multiply function
	err = client.Call("MyClass.Mul", args, &reply)
	checkError(err)

	// Call SayHello function
	name := Person{"Farrux"}
	var SayHello string
	err = client.Call("MyClass.SayHello", name, &SayHello)
	checkError(err)

	fmt.Printf(SayHello + "\n")
	fmt.Printf("Result: %d*%d=%d\n", args.A, args.B, reply)

}
