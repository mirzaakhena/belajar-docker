package remoting

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type RemoteListener struct {
	handler any
	port    int
}

func NewRemoteListener(Port int) *RemoteListener {
	return &RemoteListener{
		port: Port,
	}
}

func (r *RemoteListener) SetHandler(Handler any) {
	r.handler = Handler
}

func (r *RemoteListener) Run() {

	err := rpc.Register(r.handler)
	if err != nil {
		log.Fatal(err.Error())
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", r.port))
	if err != nil {
		log.Fatal("listen error:", err)
	}

	fmt.Printf("\nrpc server is running and waiting for rpc request from client...\n\n")

	err = http.Serve(listener, nil)
	if err != nil {
		return
	}

}
