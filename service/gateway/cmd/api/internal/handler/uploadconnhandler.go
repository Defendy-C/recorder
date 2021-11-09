package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"recorder/service/gateway/cmd/api/internal/logic"
	"recorder/service/gateway/cmd/api/internal/svc"
	"recorder/service/gateway/cmd/api/internal/types"
)

func uploadConnHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadConnReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadConnLogic(r.Context(), ctx)
		resp, err := l.UploadConn(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
