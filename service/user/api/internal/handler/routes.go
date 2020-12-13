// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "hey-go-zero/service/user/api/internal/handler/auth"
	noauth "hey-go-zero/service/user/api/internal/handler/noauth"
	"hey-go-zero/service/user/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/register",
				Handler: noauth.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login",
				Handler: noauth.LoginHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserCheck},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/user/info/self",
					Handler: auth.UserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/user/info/edit",
					Handler: auth.UserInfoEditHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
