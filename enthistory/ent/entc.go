//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ariga/edit-twitter-example-app/enthistory"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{}, entc.Extensions(enthistory.NewExtension()))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
