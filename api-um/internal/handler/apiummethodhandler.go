package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-aula-4/api-um/internal/logic"
	"go-zero-aula-4/api-um/internal/svc"
	"go-zero-aula-4/api-um/internal/types"
)

func ApiUmMethodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApiUmMethodRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiUmMethodLogic(r.Context(), svcCtx)
		resp, err := l.ApiUmMethod(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
