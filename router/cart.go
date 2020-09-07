package router

import (
	"github.com/gin-gonic/gin"
	"goshop/front-api/controller"
	"goshop/front-api/pkg/core/routerhelper"
	"goshop/front-api/pkg/middleware"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("cart", new(controller.Cart), r, middleware.VerifyToken())
		g.Get("/index")
		g.Post("/add")
		g.Get("/delete")
	})
}
