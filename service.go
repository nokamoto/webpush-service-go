package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	webpush "github.com/nokamoto/webpush-go"
	ptype "github.com/nokamoto/webpush-go/types/webpush/protobuf"
	pb "github.com/nokamoto/webpush-service-go/grpc/webpush/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type service struct {
	client *webpush.Client
}

func (s *service) Send(_ context.Context, msg *pb.Message) (*empty.Empty, error) {
	// this marshal/unmarshal is redundant, but needed.
	buf, err := proto.Marshal(msg)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}
	pmsg := ptype.Message{}
	err = proto.Unmarshal(buf, &pmsg)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	res, err := s.client.Send(pmsg)
	if err != nil {
		return nil, grpc.Errorf(codes.Unimplemented, err.Error())
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 201:
		return &empty.Empty{}, nil
	default:
	}

	return nil, grpc.Errorf(codes.Unimplemented, "not implemented yet")
}
