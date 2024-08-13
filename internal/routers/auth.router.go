package routers

import (
	"backend_coffeeShop.go/internal/handlers"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func authRouter(g *gin.Engine , d * sqlx.DB) {
	router := g.Group("/user")

	var userRepo repository.UserRepositoryInterface = repository.NewUserRepository(d)
	var authRepo repository.AuthRepositoryInterface = repository.NewAuthRepository(d)
	handler := handlers.NewAuthHandler(userRepo , authRepo)
	
	router.POST("/register" , handler.Register)
	router.POST("/login" , handler.Login)
	router.GET("/" , middleware.AuthJwtMiddleware(), handler.FetchAllUser)
	router.GET("/:id", middleware.AuthJwtMiddleware(),handler.FetchDetailUser)
	router.PATCH("/:id",handler.Update)	
	router.DELETE("/:id",handler.Delete)		
}
