// cmd/http/main.go (or wherever your main is)
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

	// Tailscale DSN — '@' in password encoded as %40
	connStr := "postgres://m56oo:1010204080%40Ph@100.115.176.71:5432/devdb?sslmode=require"
	db := mydb.MustOpen(ctx, connStr) // returns *pgxpool.Pool
	defer db.Close()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(mw.CORS(), mw.ResponseTimeGin(), mw.SecurityHeaderGin(), mw.XssValidatorGin())

	// mount routes
	router.ManufacturerRouter(r, db)

	log.Println("🚀 Server listening on :8001")
	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
