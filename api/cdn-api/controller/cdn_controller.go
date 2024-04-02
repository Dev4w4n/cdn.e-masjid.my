package controller

import (
	"net/http"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/service"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/cdn-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CDNController struct {
	engine     *gin.Engine
	cdnService service.CDNService
	env        utils.Environment
}

func NewCDNController(engine *gin.Engine, service service.CDNService, env *utils.Environment) *CDNController {
	controler := &CDNController{
		engine:     engine,
		cdnService: service,
		env:        *env,
	}

	controler.engine.POST(env.DeployURL+"/upload", controler.Upload)

	return controler
}

func (controller *CDNController) Upload(ctx *gin.Context) {
	log.Info().Msg("Saving file")

	request := model.Request{}
	err := ctx.ShouldBindJSON(&request)
	utils.WebError(ctx, err, "failed to bind JSON")

	response, err := controller.cdnService.Upload(request, &controller.env)
	utils.WebError(ctx, err, "failed to save image")

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}
