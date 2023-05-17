package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	//mage:import data
	_ "github.com/nherson/psc/api/magefiles/data"
	//mage:import fightodds
	_ "github.com/nherson/psc/api/magefiles/fightodds"
)

type Ent mg.Namespace

func (Ent) Schema(name string) error {
	return sh.Run("go", "run", "-mod=mod", "entgo.io/ent/cmd/ent", "new", "--target", "api/ent/schema", name)
}

func (Ent) Generate() error {
	fmt.Println("Generating Ent models")
	return sh.Run("go", "generate", "./api/ent")
}

func (Ent) Diff(name string) error {
	fmt.Println("Generating diff for SQL migration")
	return sh.Run("atlas", "migrate", "diff", name, "--dir", "file://api/ent/migrate/migrations", "--to", "ent://api/ent/schema", "--dev-url", "docker://postgres/15/test?search_path=public")
}

func (Ent) Migrate() error {
	return sh.Run("atlas", "migrate", "apply", "--dir", "file://api/ent/migrate/migrations", "--url", connectionString())
}

func (Ent) Shell() error {
	_, err := sh.Exec(map[string]string{}, os.Stdout, os.Stderr, "cockroach", "sql", "--url", connectionString())
	return err
}

type Proto mg.Namespace

func (Proto) Generate() error {
	err := sh.Run("buf", "generate")
	if err != nil {
		return err
	}

	return sh.Run("mv", "api/proto/api/v1/psc_connect.ts", "api/proto/api/v1/psc_pb.ts", "web/src/api/")
}

func (Proto) Format() error {
	return sh.Run("gofmt", "-w", "api/internal/proto/api/v1")
}

type Web mg.Namespace

func (Web) Start() error {
	_, err := sh.Exec(nil, os.Stdout, os.Stderr, "yarn", "--cwd", "web", "start")
	return err
}

func Deploy() error {
	return sh.Run("flyctl", "deploy")
}

func connectionString() string {
	err := godotenv.Load()
	if err != nil {
		panic("could not pull db creds from .env file")
	}

	dbHost := os.Getenv("COCKROACH_DB_HOST")
	dbUser := os.Getenv("COCKROACH_DB_USER")
	dbPassword := os.Getenv("COCKROACH_DB_PASSWORD")
	dbName := os.Getenv("COCKROACH_DB_NAME")

	if dbHost == "" {
		panic("COCKROACH_DB_HOST not set in .env file")
	}
	if dbUser == "" {
		panic("COCKROACH_DB_USER not set in .env file")
	}
	if dbPassword == "" {
		panic("COCKROACH_DB_PASSWORD not set in .env file")
	}
	if dbName == "" {
		panic("COCKROACH_DB_NAME not set in .env file")
	}

	return fmt.Sprintf("postgresql://%s:%s@%s:26257/%s?sslmode=verify-full", dbUser, dbPassword, dbHost, dbName)
}
