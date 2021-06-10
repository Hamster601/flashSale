package api

import (
	"net/http"

	"github.com/Hamster601/flashSale/application/domain/event"
	"github.com/Hamster601/flashSale/application/domain/shop"
	"github.com/Hamster601/flashSale/application/domain/stock"
	auth "github.com/Hamster601/flashSale/application/domain/user"
	"github.com/Hamster601/flashSale/application/infrastructures/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Event struct{}

type Shop struct{}

func (e *Event) List(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: event.TestData,
		Msg:  "ok",
	}
	status := http.StatusOK

	//logrus.Info("event list")

	ctx.JSON(status, resp)
}

func (e *Event) Info(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK

	logrus.Info("event info")

	ctx.JSON(status, resp)
}

func (e *Event) Subscribe(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK

	logrus.Info("event subscribe")

	ctx.JSON(status, resp)
}

func (s *Shop) AddCart(ctx *gin.Context) {
	resp := &utils.Response{
		Code: 0,
		Data: nil,
		Msg:  "ok",
	}
	status := http.StatusOK

	params := struct {
		GoodsID string `json:"goods_id"`
		EventID string `json:"event_id"`
	}{}
	var userInfo *auth.Info
	if v, ok := ctx.Get("userInfo"); ok {
		userInfo, _ = v.(*auth.Info)
	}

	err := ctx.BindJSON(&params)
	if err != nil || params.EventID == "" || params.GoodsID == "" || userInfo == nil {
		resp.Msg = "bad request"
		status = http.StatusBadRequest
		ctx.JSON(status, resp)
		return
	}
	logrus.Info(params)

	st, _ := stock.NewMemStock(params.EventID, params.GoodsID)
	if s, _ := st.Sub(userInfo.UID); s < 0 {
		resp.Code = shop.ErrNoStock
		resp.Msg = "no stock"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	conn, w, err1 := ctx.Writer.Hijack()
	if err1 != nil {
		resp.Msg = "bad request"
		status = http.StatusBadRequest
		ctx.JSON(status, resp)
		return
	}
	logrus.Info("shop add cart")
	shopCtx := &shop.Context{
		Request: ctx.Request,
		Conn:    conn,
		Writer:  w,
		GoodsID: params.GoodsID,
		EventID: params.EventID,
		UID:     userInfo.UID,
	}
	shop.Handle(shopCtx)
}

type User struct{}

func (u User) Login(ctx *gin.Context) {
	var (
		uid    string
		passwd string
		ok     bool
	)
	if uid, ok = ctx.GetPostForm("uid"); !ok {
		utils.Abort(ctx, http.StatusUnauthorized, "login failed")
		return
	}
	if passwd, ok = ctx.GetPostForm("password"); !ok {
		utils.Abort(ctx, http.StatusUnauthorized, "login failed")
		return
	}
	info, token := auth.Login(uid, passwd)
	if info != nil {
		ctx.Header(auth.TokenHeader, auth.TokenPrefix+token)
		utils.ResponseJSON(ctx, http.StatusOK, "success", nil)
	} else {
		utils.Abort(ctx, http.StatusUnauthorized, "login failed")
	}
}
