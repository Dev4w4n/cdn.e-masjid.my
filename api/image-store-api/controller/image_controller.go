package controller

import (
	"net/http"

	"github.com/Dev4w4n/cdn.e-masjid.my/api/image-store-api/model"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/image-store-api/repository"
	"github.com/Dev4w4n/cdn.e-masjid.my/api/image-store-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ImageController struct {
	engine          *gin.Engine
	imageRepository repository.ImageRepository
	env             utils.Environment
}

func NewImageController(engine *gin.Engine, repo repository.ImageRepository, env *utils.Environment) *ImageController {
	controler := &ImageController{
		engine:          engine,
		imageRepository: repo,
		env:             *env,
	}

	controler.engine.POST(env.DeployURL+"/saveImage", controler.SaveImage)

	return controler
}

func (controller *ImageController) SaveImage(ctx *gin.Context) {
	log.Info().Msg("Saving image")

	saveImage := model.Request{}
	err := ctx.ShouldBindJSON(&saveImage)
	utils.WebError(ctx, err, "failed to bind JSON")

	response, err := controller.imageRepository.SaveImage(saveImage, &controller.env)
	utils.WebError(ctx, err, "failed to save image")

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, response)
}
