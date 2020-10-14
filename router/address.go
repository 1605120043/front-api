package router

import (
	"github.com/gin-gonic/gin"
	"goshop/front-api/controller"
	"goshop/front-api/pkg/core/routerhelper"
	"goshop/front-api/pkg/middleware"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("address", new(controller.Address), r, middleware.VerifyToken())
		g.Get("/index")
		g.Get("/detail")
		g.Post("/add")
		g.Post("/edit")
		g.Post("/delete")
	})
}
