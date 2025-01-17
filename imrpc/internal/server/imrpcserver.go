// Code generated by goctl. DO NOT EDIT.
// Source: imrpc.proto

package server

import (
	"context"

	"im/imrpc/imrpc"
	"im/imrpc/internal/logic"
	"im/imrpc/internal/svc"
)

type ImrpcServer struct {
	svcCtx *svc.ServiceContext
	imrpc.UnimplementedImrpcServer
}

func NewImrpcServer(svcCtx *svc.ServiceContext) *ImrpcServer {
	return &ImrpcServer{
		svcCtx: svcCtx,
	}
}

func (s *ImrpcServer) Ping(ctx context.Context, in *imrpc.Request) (*imrpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
