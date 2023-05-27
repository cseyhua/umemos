package service

import (
	"cseyhua/memos/internal/api"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func (service Service) registryAuth(g *gin.RouterGroup) {
	g.GET("ping", ping)
	g.GET("signin", signin)
}

func ping(ctx *gin.Context) {
	ctx.SecureJSON(200, "服务可达!")
}

func signin(ctx *gin.Context) {
	_signin := &api.SignIn{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(_signin); err != nil {
		ctx.SecureJSON(404, "获取用户信息失败")
		return
	}
	ctx.SecureJSON(200, "登录服务")
}
