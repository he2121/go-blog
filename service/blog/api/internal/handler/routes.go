// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	blog "github.com/he2121/go-blog/service/blog/api/internal/handler/blog"
	"github.com/he2121/go-blog/service/blog/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/blogs",
				Handler: blog.GetBlogListHandler(serverCtx),
			},
		},
	)
}
