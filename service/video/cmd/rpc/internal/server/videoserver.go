// Code generated by goctl. DO NOT EDIT!
// Source: video.proto

package server

import (
	"context"

	"recorder/service/video/cmd/rpc/internal/logic"
	"recorder/service/video/cmd/rpc/internal/svc"
	"recorder/service/video/cmd/rpc/video"
)

type VideoServer struct {
	svcCtx *svc.ServiceContext
}

func NewVideoServer(svcCtx *svc.ServiceContext) *VideoServer {
	return &VideoServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServer) Create(ctx context.Context, in *video.CreateReq) (*video.CreateResp, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *VideoServer) GetVideoList(ctx context.Context, in *video.GetVideoListReq) (*video.GetVideoListResp, error) {
	l := logic.NewGetVideoListLogic(ctx, s.svcCtx)
	return l.GetVideoList(in)
}

func (s *VideoServer) GetOne(ctx context.Context, in *video.GetOneReq) (*video.GetOneResp, error) {
	l := logic.NewGetOneLogic(ctx, s.svcCtx)
	return l.GetOne(in)
}

func (s *VideoServer) Delete(ctx context.Context, in *video.DeleteReq) (*video.DeleteResp, error) {
	l := logic.NewDeleteLogic(ctx, s.svcCtx)
	return l.Delete(in)
}
