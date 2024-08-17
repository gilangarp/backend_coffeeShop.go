package handlers

import (
	"backend_coffeeShop.go/internal/models"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/pkg"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type PromoHandler struct {
	repository.PromoRepositoryInterface
}

func NewPromoHandler(r repository.PromoRepositoryInterface) *PromoHandler {
	return &PromoHandler{r}
}

func(h *PromoHandler) Create(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.Promo{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	if _, err := govalidator.ValidateStruct(&body); err != nil {
        response.BadRequest("create data failed", err.Error())
        return
    }

	result, err := h.CreateData(&body)
    if err != nil {
        response.InternalServerError("create data failed", err.Error())
        return
    }

    response.Success("create data success", result)
}

func(h *PromoHandler) FetchAll(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	result, err := h.GetData()
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}

func(h *PromoHandler) Update(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.Promo{}
	id := ctx.Param("id")

	if err := ctx.ShouldBind(&body); err != nil {
		response.InternalServerError("update data failed", err.Error())
		return
	}

	_ ,err := govalidator.ValidateStruct(&body)
	if err != nil {
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

func(h *PromoHandler) Delete(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	result, err := h.DeleteData(id)
	if err != nil {
		response.InternalServerError("Delete data failed", err.Error())
		return
	}

	response.Success("Delete data success", result)
}