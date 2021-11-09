package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"recorder/service/gateway/cmd/api/internal/logic"
	"recorder/service/gateway/cmd/api/internal/svc"
	"recorder/service/gateway/cmd/api/internal/types"
)

func uploadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUploadLogic(r.Context(), ctx)

		resp, err := l.Upload(req, r.Body)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
