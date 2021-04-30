package logic

import (
	"context"

	"rpc/internal/svc"
	"rpc/proto"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *proto.IdRequest) (*proto.UserResponse, error) {
	// todo: add your logic here and delete this line

	return &proto.UserResponse{}, nil
}
