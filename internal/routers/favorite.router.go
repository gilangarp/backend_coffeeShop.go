package routers

import (
	"backend_coffeeShop.go/internal/handlers"
	"backend_coffeeShop.go/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, d *sqlx.DB) {
	route := g.Group("/favorite")

	var r repository.FavoriteRepositoryInterface = repository.NewFavoriteRepository(d)
	handler := handlers.NewFavoriteHandler(r)

	route.POST("/:id", handler.Post)
	route.GET("/:id", handler.FetchAll)
	route.DELETE("/:id" ,handler.Delet)
}