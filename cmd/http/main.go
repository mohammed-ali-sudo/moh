package main

import (
	"context"
	"log"
	router "moh/internal/adapters/http/router"
	mydb "moh/shared/db"

	mw "moh/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	connStr := "postgres://postgres:1010204080@database-postgressdb.clgkywaycm2o.eu-north-1.rds.amazonaws.com:5432/pharmasto?sslmode=require"

	db := mydb.MustOpen(ctx, connStr) // returns *pgxpool.Pool
	defer db.Close()

	// Gin engine
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// apply middlewares (or replace with your own from shared/middlewares)
	r.Use(mw.CORS(), mw.ResponseTimeGin(), mw.SecurityHeaderGin(), mw.XssValidatorGin())

	// Mount /manufacturer routes
	router.ManufacturerRouter(r, db)

	log.Println("🚀 Server listening on :8001")
	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
