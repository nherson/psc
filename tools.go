//go:build tools

package main

import (
	_ "entgo.io/ent"
	_ "entgo.io/ent/cmd/ent"
	_ "entgo.io/ent/entc/gen"
	_ "github.com/twitchtv/twirp/protoc-gen-twirp"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
