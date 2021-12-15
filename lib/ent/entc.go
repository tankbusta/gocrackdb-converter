//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate(
		"./lib/ent/schema",
		&gen.Config{
			Target: "./lib/ent",
			Header: "// Code generated by entc, DO NOT EDIT.",
			Features: []gen.Feature{
				{
					Name: "sql/modifier",
				},
				{
					Name: "sql/upsert",
				},
				{
					Name: "entql",
				},
			},
		},
	); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
