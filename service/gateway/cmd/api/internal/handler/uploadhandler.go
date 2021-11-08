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
		f, _, err := r.FormFile("data")
		if err != nil {
			httpx.Error(w, err)
		}

		var data []byte
		_, err = f.Read(data)
		if err != nil {
			httpx.Error(w, err)
		}

		resp, err := l.Upload(req, data)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
