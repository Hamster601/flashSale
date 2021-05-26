package middleware

import (
	"net/http"

	"github.com/Hamster601/flashSale/application/domain/user"
	"github.com/Hamster601/flashSale/application/infrastructures/utils"
	"github.com/gin-gonic/gin"
)

func Blacklist(ctx *gin.Context) {
	data, _ := ctx.Get("UserInfo")
	info, ok := data.(*user.Info)
	if !ok {
		utils.Abort(ctx, http.StatusUnauthorized, "need login")
		return
	}
	if utils.InBlacklist(info.UID) {
		utils.Abort(ctx, http.StatusForbidden, "blocked")
		return
	}
	ctx.Next()
}
