package controller

import (
	"net/http"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/upload-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/upload-api/service"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/upload-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UploadController struct {
	engine        *gin.Engine
	uploadService service.UploadService
	env           utils.Environment
}

func NewUploadController(engine *gin.Engine, service service.UploadService, env *utils.Environment) *UploadController {
	controler := &UploadController{
		engine:        engine,
		uploadService: service,
		env:           *env,
	}

	controler.engine.POST(env.DeployURL+"/upload", controler.SaveImage)

	return controler
}

func (controller *UploadController) SaveImage(ctx *gin.Context) {
	log.Info().Msg("Saving image")

	saveImage := model.Request{}
	err := ctx.ShouldBindJSON(&saveImage)
	utils.WebError(ctx, err, "failed to bind JSON")

	response, err := controller.uploadService.SaveFile(saveImage, &controller.env)
	utils.WebError(ctx, err, "failed to save image")

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}
