package routers

import (
	"backend_coffeeShop.go/internal/handlers"
	"backend_coffeeShop.go/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func authRouter(g *gin.Engine , d * sqlx.DB) {
	router := g.Group("/user")

	var userRepo repository.UserRepositoryInterface = repository.NewUserRepository(d)
	var authRepo repository.AuthRepositoryInterface = repository.NewAuthRepository(d)
	handler := handlers.NewAuthHandler(userRepo , authRepo)
	
	router.POST("/register" , handler.Register)
	router.GET("/" , handler.FetchAllUser)
	router.POST("/login" , handler.Login)
}
