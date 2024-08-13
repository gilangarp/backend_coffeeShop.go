package handlers

import (
	"fmt"
	"math/rand"

	"backend_coffeeShop.go/internal/models"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/pkg"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	repository.ProfileRepositoryInterface
	pkg.Cloudinary
}

func NewProfileHandler(r repository.ProfileRepositoryInterface , cld pkg.Cloudinary) *ProfileHandler {
	return &ProfileHandler{r , cld}
}

func (h *ProfileHandler)CreateProfile(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.Profile{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
        response.BadRequest("create data failed", err.Error())
        return
    }

	/* img */
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		response.BadRequest("create data failed, upload file failed", err.Error())
		return
	}

	mimeType := header.Header.Get("Content-Type")
	if mimeType != "image/jpg" && mimeType != "image/png" {
		response.BadRequest("create data failed, upload file failed, wrong file type", err.Error())
		return
	}

	randomNumber := rand.Int()
	fileName := fmt.Sprintf("go-profile-%d", randomNumber)
	uploadResult, err := h.UploadFile(ctx, file, fileName)
	if err != nil {
		response.BadRequest("create data failed, upload file failed", err.Error())
		return
	}
	body.Image = uploadResult.SecureURL

	result, err := h.CreatedData(&body, id)
    if err != nil {
        response.InternalServerError("create data failed", err.Error())
        return
    }

    response.Success("create data success", result)
}

func (h *ProfileHandler)FetchAllProfile(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	result, err := h.GetAllData()
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func (h *ProfileHandler)FetchDetailProfile(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")
	
	result, err := h.GetDetailData(id)
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func (h *ProfileHandler) Update(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	body := models.Profile{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.InternalServerError("update data failed", err.Error())
		return
	}

	_ ,err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}

	id := ctx.Param("id")
	result, err := h.EditData(&body, id)
	if err != nil {
		response.InternalServerError("update data failed", err.Error())
		return
	}

	response.Success("update data success", result)
}

func (h *ProfileHandler) Delete(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	result, err := h.DeleteData(id)
	if err != nil {
		response.InternalServerError("Delete data failed", err.Error())
		return
	}

	response.Success("Delete data success", result)
}