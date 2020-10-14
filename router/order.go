package router

import (
	"goshop/front-api/controller"
	"goshop/front-api/pkg/core/routerhelper"
	"goshop/front-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("order", new(controller.Order), r, middleware.VerifyToken())
		g.Get("/index")
		g.Get("/info")
		g.Get("/get-user-order-status-count", "GetUserOrderStatusCount")
		g.Post("/create-order", "CreateOrder")
		g.Post("/cancel-order", "CancelOrder")
		g.Post("/delete-order", "DeleteOrder")
	})
}
