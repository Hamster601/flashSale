package server

import (
	"net"

	"github.com/Hamster601/flashSale/application/domain/shop"
	"github.com/Hamster601/flashSale/application/infrastructures/pkg/redis"
	"github.com/Hamster601/flashSale/application/infrastructures/utils"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var lis net.Listener

func Run() error {
	var err error
	bind := viper.GetString("api.bind")
	logrus.Info("run api server on ", bind)
	lis, err = utils.Listen("tcp", bind)
	if err != nil {
		return err
	}

	g := gin.New()

	// 更新程序，给老版本发送信号
	go utils.UpdateProc("api")

	if err := redis.Init(); err != nil {
		return err
	}
	// 监控黑名单变更
	utils.WatchBlacklist()
	// 初始化路由
	initRouters(g)
	// 初始化 shop
	shop.Init()
	pprof.Register(g)
	// 运行服务
	return g.RunListener(lis)
}

func Exit() {
	lis.Close()
	// TODO: 等待请求处理完
	// time.Sleep(10 * time.Second)
	logrus.Info("api server exit")
}
