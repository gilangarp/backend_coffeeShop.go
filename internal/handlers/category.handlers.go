package handlers

import (
	"fmt"

	"backend_coffeeShop.go/internal/models"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/pkg"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	repository.CategoryRepositoryInterface
}

func NewCategoryHandler(r repository.CategoryRepositoryInterface) *CategoryHandler {
	return &CategoryHandler{r}
}

func (h *CategoryHandler)CreatedCategory(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.Category{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed,", err.Error())
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
        response.BadRequest("create data failed", err.Error())
        return
    }

	fmt.Printf("ini body:%s" , body)
	result, err := h.CreatedData(&body)
    if err != nil {
        response.InternalServerError("create data failed", err.Error())
        return
    }

    response.Success("create data success", result)
}

func (h *CategoryHandler)FetchAllCategory(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	result, err := h.GetData()
    if err != nil {
        response.InternalServerError("Get data failed", err.Error())
        return
    }

    response.Success("Get data success", result)
}

func (h *CategoryHandler)Update(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.Category{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("update data failed", err.Error())
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
        response.BadRequest("update data failed", err.Error())
        return
    }

	result, err := h.UpdateData(&body, id)
    if err != nil {
        response.InternalServerError("update data failed", err.Error())
        return
    }

    response.Success("update data success", result)
}

func (h *CategoryHandler)Delete(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	result, err := h.DeleteData(id)
	if err != nil {
		response.InternalServerError("Delete data failed", err.Error())
		return
	}

	response.Success("Delete data success", result)
}