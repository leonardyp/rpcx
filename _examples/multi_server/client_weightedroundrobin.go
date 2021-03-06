package main

import (
	"fmt"
	"time"

	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/clientselector"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	server1 := &clientselector.ServerPeer{Network: "tcp", Address: "127.0.0.1:8972", Weight: 9}
	server2 := &clientselector.ServerPeer{Network: "tcp", Address: "127.0.0.1:8973", Weight: 1}

	servers := []*clientselector.ServerPeer{server1, server2}

	s := clientselector.NewMultiClientSelector(servers, rpcx.WeightedRoundRobin, 10*time.Second)

	for i := 0; i < 10; i++ {
		callServer(s)
	}
}

func callServer(s rpcx.ClientSelector) {
	client := rpcx.NewClient(s)
	args := &Args{7, 8}
	var reply Reply
	err := client.Call("Arith.Mul", args, &reply)
	if err != nil {
		fmt.Printf("error for Arith: %d*%d, %v \n", args.A, args.B, err)
	} else {
		fmt.Printf("Arith: %d*%d=%d \n", args.A, args.B, reply.C)
	}

	client.Close()
}
