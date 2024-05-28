package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-aula-4/api-dois/internal/logic"
	"go-zero-aula-4/api-dois/internal/svc"
	"go-zero-aula-4/api-dois/internal/types"
)

func ApiDoisMethodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApiDoisMethodRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiDoisMethodLogic(r.Context(), svcCtx)
		resp, err := l.ApiDoisMethod(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
