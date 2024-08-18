package routers

import (
	"backend_coffeeShop.go/internal/handlers"
	"backend_coffeeShop.go/internal/repository"
	"backend_coffeeShop.go/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func productRouter(g *gin.Engine, d *sqlx.DB) {
	router := g.Group("/product")

	var repo repository.ProductRepositoryInterface = repository.NewProductRepository(d)
	var cld pkg.Cloudinary = *pkg.NewCloudinaryUtil()
	handler := handlers.NewProductHandler(repo, cld)

	router.POST("/", handler.Post)
	router.GET("/", handler.FetchAll)
	router.GET("/:id", handler.FetchDetail)
	router.PATCH("/:id" , handler.Update)
	router.DELETE("/:id", handler.Delete)


}