package gift

import (
	"strconv"
	"ucenter/app/controllers"
	"ucenter/app/logs"
	"ucenter/models"

	"github.com/gin-gonic/gin"
)

//获取礼物列表
func GiftList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	r, err := models.GetGiftList(page, limit)
	if err != nil {
		logs.Logger.Error(err)
	}
	controllers.SuccessStr(c, r, "")
}

//送礼
func Send(c *gin.Context) {
	rs, _ := c.Get("_user")
	user, _ := rs.(*models.UserModel)

	giftId, _ := strconv.Atoi(c.PostForm("gift_id"))
	toUserId, _ := strconv.Atoi(c.PostForm("uid"))
	toUserId64 := int64(toUserId)

	if user.Id == toUserId64 {
		controllers.ErrorNoData(c, "Can't give yourself a gift")
		return
	}

	gift, err := models.GetGiftRow(giftId)
	if err != nil {
		logs.Logger.Error(err)
		controllers.ErrorNoData(c, "Gift not found")
		return
	}

	toUser := models.GetUser(toUserId64, "", "", "")
	if toUser == nil || toUser.Id < 1 {
		controllers.ErrorNoData(c, "Non-existent users")
		return
	}

	balance := user.GetUserBalance()
	if balance < 1 || balance < gift.Bi {
		controllers.ErrorNoData(c, "Insufficient balance")
		return
	}

	err = models.AddGiftHis(gift, user.Id, toUserId64)
	if err != nil {
		controllers.ErrorNoData(c, "Failed gift giving")
		return
	}
	controllers.SuccessStr(c, nil, "")
}
