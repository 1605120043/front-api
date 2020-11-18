package router

import (
	"github.com/gin-gonic/gin"
	"goshop/front-api/controller"
	"goshop/front-api/pkg/core/routerhelper"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("common", new(controller.Common), r)
		g.Get("/get-area-list")
		g.Post("/mobile-login")
		g.Post("/send-code")
		g.Post("/get-wx-openid")
		g.Post("/wx-login")
	})
}
