package handler

import (
	"net/http"

	"github.com/he2121/go-blog/service/blog/api/internal/logic/blog"
	"github.com/he2121/go-blog/service/blog/api/internal/svc"
	"github.com/he2121/go-blog/service/blog/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetBlogListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBlogListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetBlogListLogic(r.Context(), ctx)
		resp, err := l.GetBlogList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
