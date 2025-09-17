// cmd/http/main.go (or wherever your main is)
package main

import (
	"context"
	"log"

	router "moh/internal/adapters/http/router"
	mydb "moh/shared/db"
	mw "moh/shared/middlewares"

	execs "moh/internal/adapters/grpc"

	"github.com/gin-gonic/gin"
)

func main() {
	// grpc client

	ctx := context.Background()
	const execsAddr = "127.0.0.1:50051" // gRPC server address
	const connStr = "postgres://m56oo:1010204080%40Ph@100.115.176.71:5432/devdb?sslmode=require"

	// --- hard-coded configs (change as needed) ---

	// gRPC client to the other microservice
	ExClient, err := execs.New(execsAddr)
	if err != nil {
		log.Fatal("connect execs:", err)
	}
	defer ExClient.Close()
	// grpc cleint end

	// Tailscale DSN — '@' in password encoded as %40
	db := mydb.MustOpen(ctx, connStr) // returns *pgxpool.Pool
	defer db.Close()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(mw.CORS(), mw.ResponseTimeGin(), mw.SecurityHeaderGin(), mw.XssValidatorGin())

	// mount routes
	router.ManufacturerRouter(r, db, ExClient)

	log.Println("🚀 Server listening on :8001")
	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
