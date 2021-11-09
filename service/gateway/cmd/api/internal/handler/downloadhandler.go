package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"recorder/service/gateway/cmd/api/internal/logic"
	"recorder/service/gateway/cmd/api/internal/svc"
	"recorder/service/gateway/cmd/api/internal/types"
)

func downloadHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDownloadLogic(r.Context(), ctx)
		err := l.Download(req, w)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
