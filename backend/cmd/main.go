package main

import (
	"backend/db"
	"backend/internal/user"
	"backend/internal/ws"
	"backend/router"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetSession())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	r := gin.Default()

	r.Use(cors.Default())

	router.InitRouter(userHandler, wsHandler)
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
