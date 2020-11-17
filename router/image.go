package router

import (
	"goshop/front-api/controller"
	"goshop/front-api/pkg/core/routerhelper"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("image", new(controller.Image), r)
		g.Get("/get-image")
		g.Post("/upload")
	})
}
