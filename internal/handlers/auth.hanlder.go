package handlers

import (
	"backend_coffeeShop.go/internal/models"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/pkg"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repository.UserRepositoryInterface
	repository.AuthRepositoryInterface
}

func NewAuthHandler(userRepo repository.UserRepositoryInterface , authRepo repository.AuthRepositoryInterface ) *AuthHandler{
	return &AuthHandler{userRepo,authRepo}
}

func (h *AuthHandler) Register(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.User{}

	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	_ ,err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	body.Password , err = pkg.HashPassword(body.Password)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}

	result , err := h.CreateData(&body)
	if err != nil {
		response.BadRequest("create data failed", err.Error())
		return
	}
	response.Created("create data success", result)
}

func (h *AuthHandler) Login(ctx *gin.Context){
	response := pkg.NewResponse(ctx)
	body := models.UserLogin{}
	
	if err := ctx.ShouldBind(&body); err != nil {
		response.BadRequest("Login failed", err.Error())
		return
	}

	_ ,err := govalidator.ValidateStruct(&body)
	if err != nil {
		response.BadRequest("Login failed", err.Error())
		return
	}

	result, err := h.GetByEmail(body.Email)
	if err != nil {
		response.BadRequest("Login failed", err.Error())
		return
	}
	err = pkg.VerifyPassword(result.Password, body.Password)
	if err != nil {
		response.Unauthorized("wrong password", err.Error())
		return
	}

	jwt := pkg.NewJWT(result.ID , result.Email )
	token , err := jwt.GenerateToken()
	if err != nil {
		response.Unauthorized("failed generate token", err.Error())
		return
	}
	
	response.Success("Login successful", map[string]interface{}{
		"token": token,
		"id":    result.ID,
	})
}

func (h *AuthHandler) FetchAllUser(ctx *gin.Context) {
	response := pkg.NewResponse(ctx)
	result, err := h.GetAllData()
	if err != nil {
		response.InternalServerError("get data failed", err.Error())
		return
	}

	response.Success("get data success", result)
}