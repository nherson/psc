package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/internal/clients/db"
	"github.com/nherson/psc/api/internal/services/data"
	"github.com/nherson/psc/api/internal/services/psc"
	"github.com/nherson/psc/api/proto/api/v1/apiv1connect"
	"github.com/nherson/psc/web"
)

var local bool

func init() {
	flag.BoolVar(&local, "local", false, "If set, will assume running locally and handle setup accordingly")
	flag.Parse()
}

func main() {
	var dbClient *ent.Client
	if local {
		err := godotenv.Load()
		if err != nil {
			panic("could not read env file")
		}
		dbClient = db.MustFromEnvWithPublicDNS()
	} else {
		dbClient = db.MustFromEnvWithPublicDNS()

	}

	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	pscServer := &psc.PSCServer{
		DB: dbClient,
	}

	dataServer := &data.Service{
		DB: dbClient,
	}

	apiMux.Handle(apiv1connect.NewPSCServiceHandler(pscServer))
	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	mux.HandleFunc("/data/upcoming", dataServer.PublicCSV)
	mux.HandleFunc("/data/upcoming/nick", dataServer.NickCSV)

	mux.Handle("/", web.Handler())

	c := cors.New(cors.Options{
		AllowedHeaders:     []string{"Connect-Protocol-Version", "Content-Type"},
		OptionsPassthrough: false,
		Debug:              true,
	})

	fmt.Println("Listening on port :8080")

	http.ListenAndServe(":8080", c.Handler(mux))
}
