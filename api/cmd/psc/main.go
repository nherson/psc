package main

import (
	"net/http"

	"github.com/nherson/psc/api/internal/services/psc"
	"github.com/nherson/psc/api/proto/api/v1/apiv1connect"
)

func main() {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	pscServer := &psc.PSCServer{}
	apiMux.Handle(apiv1connect.NewPSCServiceHandler(pscServer))
	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	http.ListenAndServe(":8080", mux)
}
