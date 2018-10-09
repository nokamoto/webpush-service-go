package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	webpush "github.com/nokamoto/webpush-go"
	pb "github.com/nokamoto/webpush-service-go/grpc/webpush/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

var (
	private = flag.String("private", "AJFotoB4FS7IX6tbm5t0SGyISTQ6l54mMzpfYipdOD+N", "VAPID application server private key (base64.StdEncoding)")
	public  = flag.String("public", "BNuvjW90TpDawYyxhvK79QVyNEplaSQZOWo1CwXDmWwfya6qnyBvIx3tFvKEBetExvil4rNNRL0/ZR2WLjGEAbQ=", "VAPID application server public key (base64.StdEncoding)")
	subject = flag.String("subject", "", "mailto: or https: URI")
	port    = flag.Int("port", 9090, "grpc server port")
)

func main() {
	flag.Parse()

	priv, err := base64.StdEncoding.DecodeString(*private)
	if err != nil {
		panic(fmt.Sprintf("bad private %s - %v", *private, err))
	}

	pub, err := base64.StdEncoding.DecodeString(*public)
	if err != nil {
		panic(fmt.Sprintf("bad public %s - %v", *public, err))
	}

	if len(*subject) == 0 {
		subject = nil
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(fmt.Sprintf("listen tcp port (%d) - %v", *port, err))
	}

	fmt.Printf("listen tcp port (%d)\n", *port)

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	service := &service{client: webpush.NewClient(&http.Client{}, priv, pub, subject)}

	pb.RegisterPushServiceServer(server, service)
	reflection.Register(server)

	err = server.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("serve %v - %v", lis, err))
	}
}
