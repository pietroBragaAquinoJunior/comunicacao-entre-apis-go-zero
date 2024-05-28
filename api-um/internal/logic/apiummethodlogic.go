package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpc"
	"go-zero-aula-4/api-um/internal/svc"
	"go-zero-aula-4/api-um/internal/types"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiUmMethodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiUmMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiUmMethodLogic {
	return &ApiUmMethodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ApiDoisMethodResponse struct {
	Resposta string `json:"resposta"`
}

func (l *ApiUmMethodLogic) ApiUmMethod(req *types.ApiUmMethodRequest) (resp *types.ApiUmMethodResponse, err error) {

	respostaSincrona, err := httpc.Do(l.ctx, http.MethodGet, "http://localhost:8889/method", nil)
	if err != nil {
		fmt.Println(err)
		return &types.ApiUmMethodResponse{Resposta: "erro ao fazer a requisição"}, err
	}
	defer respostaSincrona.Body.Close()

	// Leia o corpo da resposta
	bodyBytes, err := io.ReadAll(respostaSincrona.Body)
	if err != nil {
		fmt.Println(err)
		return &types.ApiUmMethodResponse{Resposta: "erro ao ler o corpo da resposta"}, err
	}

	// Deserializa o JSON da resposta
	var resposta ApiDoisMethodResponse
	err = json.Unmarshal(bodyBytes, &resposta)
	if err != nil {
		fmt.Println(err)
		return &types.ApiUmMethodResponse{Resposta: "erro ao deserializar o corpo da resposta"}, err
	}

	// Retorne a string lida no corpo da resposta
	return &types.ApiUmMethodResponse{Resposta: resposta.Resposta}, nil

}
