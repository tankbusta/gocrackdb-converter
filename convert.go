package gocrackdb

import (
	"context"

	"github.com/tankbusta/gocrackdb-converter/lib/converter"
	"github.com/tankbusta/gocrackdb-converter/lib/ent"
	_ "github.com/tankbusta/gocrackdb-converter/lib/ent/runtime"

	"github.com/asdine/storm"
	_ "github.com/lib/pq"
)

func MigrateFromBolt(ctx context.Context, boltPath, dbConnStr string) error {
	boltdb, err := storm.Open(boltPath)
	if err != nil {
		return err
	}
	defer boltdb.Close()

	client, err := ent.Open("postgres", dbConnStr)
	if err != nil {
		return err
	}
	defer client.Close()

	converter := converter.NewConverter(boltdb, client)
	return converter.Convert(ctx)
}
