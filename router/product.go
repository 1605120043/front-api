package router

import (
	"github.com/gin-gonic/gin"
	"goshop/front-api/controller"
	"goshop/front-api/pkg/core/routerhelper"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("product", new(controller.Product), r)
		g.Get("/tag")
		g.Get("/index")
		g.Get("/detail")
	})
}
