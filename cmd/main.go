package main

import (
	"log"

	"backend_coffeeShop.go/internal/routers"
	"backend_coffeeShop.go/pkg"
	_ "github.com/joho/godotenv/autoload"
)

func main(){
	db ,err := pkg.PostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := routers.NewRouter(db)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}