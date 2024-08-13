package routers

import (
	"backend_coffeeShop.go/internal/handlers"
	"backend_coffeeShop.go/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func profileRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/profile")

	var repo repository.ProfileRepositoryInterface = repository.NewProfileRepository(d)
	handler := handlers.NewProfileHandler(repo)

	router.POST("/:id" , handler.CreateProfile)
	router.GET("/" , handler.FetchAllProfile)
	router.GET("/:id" , handler.FetchDetailProfile)
	router.PATCH("/:id" , handler.Update)
	router.DELETE("/:id" , handler.Delete)

}