package router

import (
	"goshop/front-api/controller"
	"goshop/front-api/pkg/core/routerhelper"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("banner", new(controller.BannerAd), r)
		g.Get("/index")
	})
}
