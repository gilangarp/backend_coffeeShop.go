package handlers

import (
	"backend_coffeeShop.go/internal/models"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/pkg"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	repository.ProfileRepositoryInterface
}

func NewProfileHandler(r repository.ProfileRepositoryInterface) *ProfileHandler {
	return &ProfileHandler{r}
}

func(h *ProfileHandler)CreateProfile(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.Profile{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	_ ,err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	result, err := h.CreatedData(&body, id)
	if err != nil {
		response.InternalServerError("create data failed", err.Error())
		return
	}

	response.Success("create data success", result)
}

func(h *ProfileHandler)FetchAllProfile(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	result, err := h.GetAllData()
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func(h *ProfileHandler)FetchDetailProfile(ctx *gin.Context) {
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