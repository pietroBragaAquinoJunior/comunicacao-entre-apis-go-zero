package logic

import (
	"context"

	"go-zero-aula-4/api-dois/internal/svc"
	"go-zero-aula-4/api-dois/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiDoisMethodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiDoisMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiDoisMethodLogic {
	return &ApiDoisMethodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiDoisMethodLogic) ApiDoisMethod(req *types.ApiDoisMethodRequest) (resp *types.ApiDoisMethodResponse, err error) {
	// todo: add your logic here and delete this line

	return &types.ApiDoisMethodResponse{Resposta: "Aqui é a API DOIS (2) !"}, nil
}
