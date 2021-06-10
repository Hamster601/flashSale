package server

import (
	"github.com/Hamster601/flashSale/application/api"
	middleware "github.com/Hamster601/flashSale/application/http/middlewares"
	"github.com/Hamster601/flashSale/application/infrastructures/utils"
	"github.com/gin-gonic/gin"
)

func initRouters(g *gin.Engine) {
	g.POST("/login", api.User{}.Login)

	eventCB := utils.NewCircuitBreaker(
		utils.WithDuration(100),
		utils.WithTotalLimit(20000),
		utils.WithLatencyLimit(100),
		utils.WithFailsLimit(5),
	)
	eventCBMdw := middleware.NewCircuitBreakMiddleware(eventCB)
	event := g.Group("/event").Use(eventCBMdw, middleware.NewAuthMiddleware(false))
	eventApp := api.Event{}
	event.GET("/list", eventApp.List)
	event.GET("/info", eventApp.Info)

	subscribe := g.Group("/event/subscribe").Use(middleware.NewAuthMiddleware(true))
	subscribe.POST("/", eventApp.Subscribe)

	shopCB := utils.NewCircuitBreaker(
		utils.WithDuration(100),
		utils.WithTotalLimit(1000),
		utils.WithLatencyLimit(200),
		utils.WithFailsLimit(5),
	)
	mdws := []gin.HandlerFunc{
		middleware.NewCircuitBreakMiddleware(shopCB),
		middleware.NewAuthMiddleware(true),
		middleware.Blacklist,
	}
	shop := g.Group("/shop").Use(mdws...)
	shopApp := api.Shop{}
	shop.PUT("/cart/add", shopApp.AddCart)
}
