package main

import (
	"net/http"

	v1 "github.com/nherson/psc/api/internal/proto/api/v1"
	"github.com/nherson/psc/api/internal/services/ufcstats"
)

func main() {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	ufcStatsService := &ufcstats.Service{}
	apiMux.Handle(v1.UfcStatsPathPrefix, v1.NewUfcStatsServer(ufcStatsService))

	// API routes go under /api
	mux.Handle("/api", apiMux)

	// TODO wire some embedded files to serve a react app

	http.ListenAndServe(":8080", mux)
}
