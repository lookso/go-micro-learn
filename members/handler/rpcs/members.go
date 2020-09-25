package rpcs

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	"go-micro-learn/members/proto/members"
)

type Members struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Members) Call(ctx context.Context, req *members.Request, rsp *members.Response) error {
	log.Info("Received Members.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Members) Stream(ctx context.Context, req *members.StreamingRequest, stream members.Members_StreamStream) error {
	log.Infof("Received Members.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&members.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Members) PingPong(ctx context.Context, stream members.Members_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&members.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
