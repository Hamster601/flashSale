package api

import (
	"context"

	"github.com/Hamster601/flashSale/application/api/rpc"

	"github.com/sirupsen/logrus"
)

type EventRPCServer struct {
}

func (s *EventRPCServer) EventOnline(ctx context.Context, evt *rpc.Event) (*rpc.Response, error) {
	logrus.Info("event online ", evt)
	resp := &rpc.Response{}

	return resp, nil
}

func (s *EventRPCServer) EventOffline(ctx context.Context, evt *rpc.Event) (*rpc.Response, error) {
	logrus.Info("event offline ", evt)
	resp := &rpc.Response{}

	return resp, nil
}

func (s *EventRPCServer) TopicOnline(ctx context.Context, t *rpc.Topic) (*rpc.Response, error) {
	logrus.Info("topic online ", t)
	resp := &rpc.Response{}

	return resp, nil
}

func (s *EventRPCServer) TopicOffline(ctx context.Context, t *rpc.Topic) (*rpc.Response, error) {
	logrus.Info("topic offline ", t)
	resp := &rpc.Response{}

	return resp, nil
}
