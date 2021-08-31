package main

import (
	_ "easebase_web/routers"
	"easebase_web/utils"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"strings"
)

func main() {

	var FilterToken = func(ctx *context.Context) {
		logs.Info("current router path is ", ctx.Request.RequestURI, strings.Split(ctx.Request.RequestURI, "?")[0])
		if strings.Split(ctx.Request.RequestURI, "?")[0] != "/login" && ctx.Input.Header("Authorization") == "" && strings.Split(ctx.Request.RequestURI, "?")[0][0:4] != "/api" {
			logs.Error("无权访问(token is nil) !")
			ctx.ResponseWriter.WriteHeader(401)
			ctx.ResponseWriter.Write([]byte("无权访问!"))
		}

		if strings.Split(ctx.Request.RequestURI, "?")[0] != "/login" && ctx.Input.Header("Authorization") != "" {
			token := ctx.Input.Header("Authorization")
			logs.Info("Authorization token: ", token)
			if utils.ValidateToken(token) != nil {
				ctx.ResponseWriter.WriteHeader(402)
				ctx.ResponseWriter.Write([]byte("认证过期!"))
				logs.Info("认证过期")
			} else {
				utils.RefreshToken(token)
				logs.Info("Token已续签")
			}
		}
	}

	web.BConfig.WebConfig.TemplateLeft = "<<<"
	web.BConfig.WebConfig.TemplateRight = ">>>"
	web.InsertFilter("/*", web.BeforeRouter, FilterToken)
	web.Run()
}
