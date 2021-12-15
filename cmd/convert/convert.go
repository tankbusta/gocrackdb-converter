package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/tankbusta/gocrackdb-converter"
)

var boltDBPath string
var pgConnStr string

func init() {
	flag.StringVar(&boltDBPath, "bolt-db", "", "path to the bolt db")
	flag.StringVar(&pgConnStr, "pg-conn-str", "", "postgres conn string")
	flag.Parse()
}

func main() {
	if boltDBPath == "" || pgConnStr == "" {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	ctx := context.Background()
	signal.NotifyContext(ctx, os.Interrupt)

	if err := gocrackdb.MigrateFromBolt(ctx, boltDBPath, pgConnStr); err != nil {
		log.Fatalf("[ X ] %s", err)
	}
}
