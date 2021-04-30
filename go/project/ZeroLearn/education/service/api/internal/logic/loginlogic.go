package logic

import (
	"context"
	"errors"

	"api/internal/models"
	"api/internal/svc"
	"api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 注册登录
func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {
	// todo: add your logic here and delete this line
	user := models.User{
		OpenId:   req.OpenId,
		NickName: req.Nickname,
		Gender:   req.Gender,
		Avatar:   req.Avar,
		City:     req.City,
	}
	r := l.svcCtx.DbEngin.Take(&user)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		l.svcCtx.DbEngin.Create(&user)
		return &types.LoginReply{
			Status: "success",
			Mes:    "注册登录成功",
		}, nil
	} else {
		return &types.LoginReply{
			Status: "success",
			Mes:    "登录成功",
		}, nil
	}
}
