package main

import (
	"encoding/base64"
	"fmt"
	"github.com/nokamoto/webpush-go"
	pb "github.com/nokamoto/webpush-service-go/grpc/webpush/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	privateKey = "AJFotoB4FS7IX6tbm5t0SGyISTQ6l54mMzpfYipdOD+N"
	publicKey  = "BNuvjW90TpDawYyxhvK79QVyNEplaSQZOWo1CwXDmWwfya6qnyBvIx3tFvKEBetExvil4rNNRL0/ZR2WLjGEAbQ="
	auth       = "LsUmSxGzGt+KcuczkTfFrQ=="
	p256dh     = "BOVFfCoBB/2Sn6YZrKytKc1asM+IOXFKz6+T1NLOnrGrRXh/xJEgiJIoFBO9I6twWDAj6OYvhval8jxq8F4K0iM="
)

func decode(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(fmt.Sprintf("%s - %v", s, err))
	}
	return b
}

func Test_service_Send_ok(t *testing.T) {
	s := &service{client: webpush.NewClient(&http.Client{}, decode(privateKey), decode(publicKey), nil)}

	server := httptest.NewServer(http.HandlerFunc(func(write http.ResponseWriter, req *http.Request) {
		write.WriteHeader(201)
	}))
	defer server.Close()

	msg := &pb.Message{
		Subscription: &pb.PushSubscription{
			Endpoint: server.URL,
			Auth:     decode(auth),
			P256Dh:   decode(p256dh),
		},
	}
	_, err := s.Send(context.Background(), msg)

	if code := status.Convert(err).Code(); code != codes.OK {
		t.Errorf("expected %v but actual %v", codes.OK, code)
	}
}
