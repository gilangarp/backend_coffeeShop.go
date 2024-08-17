package routers

import (
	"backend_coffeeShop.go/internal/handlers"
	"backend_coffeeShop.go/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func promoRouter(g *gin.Engine , d *sqlx.DB) {
	router := g.Group("/promo")

	var repo repository.PromoRepositoryInterface = repository.NewPromoRepository(d)
	handler := handlers.NewPromoHandler(repo)

	router.POST("/",handler.Create)
	router.GET("/" , handler.FetchAll)
	router.PATCH("/:id" , handler.Update)
	router.DELETE("/:id" , handler.Delete)
	
}