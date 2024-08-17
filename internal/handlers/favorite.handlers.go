package handlers

import (
	"backend_coffeeShop.go/internal/models"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/pkg"
	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	repository.FavoriteRepositoryInterface
}

func NewFavoriteHandler(r repository.FavoriteRepositoryInterface) *FavoriteHandler {
	return &FavoriteHandler{r}
}


func (h *FavoriteHandler) Post(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.Favorite{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.BadRequest("create data failed,", err.Error())
		return
	}
	
	result, err := h.CreatedData(&body, id)
	if err != nil {		
		response.BadRequest("create data failed,", err.Error())
		return
	}

	response.Success("create data success", result)
}

func (h *FavoriteHandler) FetchAll(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	result , err := h.GetDetailData(id)
	if err != nil {
		response.BadRequest("Get data failed,", err.Error())
		return
	}
	
	response.Success("Get data success", result)
}

func (h *FavoriteHandler) Delet(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	id := ctx.Param("id")

	result, err := h.DeleteData(id)
	if err != nil {
		response.InternalServerError("Delete data failed", err.Error())
		return
	}

	response.Success("Delete data success", result)
}