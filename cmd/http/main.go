package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	router "moh/internal/adapters/http/router"
	mydb "moh/shared/db"
)

func main() {
	// Open DB (returns *pgxpool.Pool)
	ctx := context.Background()
	connStr := "postgres://postgres:1010204080@database-postgressdb.clgkywaycm2o.eu-north-1.rds.amazonaws.com:5432/pharmasto?sslmode=require"
	db := mydb.MustOpen(ctx, connStr)
	defer db.Close()

	// Root router (public only for now)
	mainRouter := mux.NewRouter()
	// mainRouter.Use(middleware.CORS)
	// mainRouter.Use(middleware.ResponseTimeMw)
	// mainRouter.Use(middleware.SecurityHeader)

	// Mount your router exactly as you wrote it
	mainRouter.PathPrefix("/manufacturer").Handler(
		http.StripPrefix("/manufacturer", router.ManufacturerRouter(db)),
	)

	fmt.Println("🚀 Server listening on :8001")
	log.Fatal(http.ListenAndServe(":8001", mainRouter))
}
