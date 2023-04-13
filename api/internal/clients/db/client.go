package db

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/nherson/psc/api/ent"
)

func MustFromEnv() *ent.Client {
	dbHost := os.Getenv("COCKROACH_DB_HOST")
	dbUser := os.Getenv("COCKROACH_DB_USER")
	dbPassword := os.Getenv("COCKROACH_DB_PASSWORD")
	dbName := os.Getenv("COCKROACH_DB_NAME")

	if dbHost == "" {
		panic("COCKROACH_DB_HOST not set in environment")
	}
	if dbUser == "" {
		panic("COCKROACH_DB_USER not set in environment")
	}
	if dbPassword == "" {
		panic("COCKROACH_DB_PASSWORD not set in environment")
	}
	if dbName == "" {
		panic("COCKROACH_DB_NAME not set in environment")
	}

	connectionString := fmt.Sprintf("host=%s port=26257 user=%s dbname=%s password=%s", dbHost, dbUser, dbName, dbPassword)

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	return client
}

// Manually resolve the CockroachDB IP address and use that to establish a connection.
// Not fully debugged, but A) can connect to other external network addresses, and
// B) the CockroachDB DNS uses a .cloud TLD which, who knows, maybe is under custom use
// within Flyio and therefore not usable for external lookups (pre-empted by internal DNS).
//
// WARNING: NOT THREAD SAFE. DO THIS ONCE AT APP INIT AND MOVE ON.
func MustFromEnvWithPublicDNS() *ent.Client {

	defaultResolver := net.DefaultResolver

	// do a lil switcheroo really quick
	net.DefaultResolver = &net.Resolver{
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	// make the client
	db := MustFromEnv()

	// switch back before anyone notices!
	net.DefaultResolver = defaultResolver

	return db
}
