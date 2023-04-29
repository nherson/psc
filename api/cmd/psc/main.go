package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/nherson/psc/api/internal/clients/db"
	"github.com/nherson/psc/api/internal/services/psc"
	"github.com/nherson/psc/api/proto/api/v1/apiv1connect"
	"github.com/nherson/psc/web"
)

func main() {
	dbClient := db.MustFromEnvWithPublicDNS()

	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	pscServer := &psc.PSCServer{
		DB: dbClient,
	}
	apiMux.Handle(apiv1connect.NewPSCServiceHandler(pscServer))
	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	mux.Handle("/", web.Handler())

	c := cors.New(cors.Options{
		AllowedHeaders:     []string{"Connect-Protocol-Version", "Content-Type"},
		OptionsPassthrough: false,
		Debug:              true,
	})

	http.ListenAndServe(":8080", c.Handler(mux))
}
