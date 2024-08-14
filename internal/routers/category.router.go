package routers

import (
	"backend_coffeeShop.go/internal/handlers"
	"backend_coffeeShop.go/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func categoryRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/category")

	var repo repository.CategoryRepositoryInterface = repository.NewCategoryRepository(d)
	handler := handlers.NewCategoryHandler(repo)

	router.POST("/" , handler.CreatedCategory)
	router.GET("/" , handler.FetchAllCategory)
	router.PATCH("/:id" , handler.Update)
	router.DELETE("/:id" , handler.Delete)
}