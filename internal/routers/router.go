package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:5173"},
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))
	
	authRouter(router , db)
	profileRouter(router , db)
	categoryRouter(router,db)
	productRouter(router , db)
	promoRouter(router , db)
	favorite(router , db)

	return router
}