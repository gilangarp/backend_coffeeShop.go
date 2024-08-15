package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	authRouter(router , db)
	profileRouter(router , db)
	categoryRouter(router,db)
	productRouter(router , db)

	return router
}