package info

import (
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/log"
	"github.com/AH-dark/apollo/pkg/serializer"
	"github.com/AH-dark/apollo/vo"
	"github.com/gin-gonic/gin"
)

func SiteInfoBasicHandler(c *gin.Context) {
	settings, err := model.Global.Setting.GetSettingsByType(model.SettingTypeBasic)
	if err != nil {
		log.Log().
			WithField(log.FieldGinContext, c).
			WithError(err).
			Error("Get settings error")
		c.JSON(500, serializer.NewHttpError(500, err.Error(), c.MustGet("request_id").(string)))
	}

	c.JSON(200, serializer.NewSuccessResponse(vo.BuildSettingsListVO(settings)))
}
