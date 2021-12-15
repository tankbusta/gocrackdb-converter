package converter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tankbusta/gocrackdb-converter/lib/ent"

	"github.com/asdine/storm"
)

type Converter struct {
	from *storm.DB
	to   *ent.Client
}

type ConvertFunc func(ctx context.Context, from *storm.DB, to *ent.Tx) error

type converterLibraryEntry struct {
	Name string
	Func ConvertFunc
}

var converterLibrary = []converterLibraryEntry{
	{
		Name: "base_users",
		Func: convertBaseUsers,
	},
	{
		Name: "base_tasks",
		Func: convertBaseTasks,
	},
	{
		Name: "convert_passwords",
		Func: convertPasswords,
	},
}

func NewConverter(from *storm.DB, to *ent.Client) *Converter {
	return &Converter{
		from: from,
		to:   to,
	}
}

func (s *Converter) Convert(ctx context.Context) error {
	// Run the auto migration tool.
	if err := s.to.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	var bdbVersion float32
	if err := s.from.Get("options", "version", &bdbVersion); err != nil {
		return err
	}

	if bdbVersion != 1.0 {
		return fmt.Errorf("unknown bolt version: %f", bdbVersion)
	}

	txn, err := s.to.Tx(ctx)
	if err != nil {
		return err
	}
	defer txn.Rollback()

	startedConverterAt := time.Now()
	for _, migrator := range converterLibrary {
		startedMigrationAt := time.Now()
		log.Printf("[ ! ] Starting Migration `%s`\n", migrator.Name)

		if err := migrator.Func(ctx, s.from, txn); err != nil {
			return fmt.Errorf("failed to migrate `%s`: %w", migrator.Name, err)
		}

		log.Printf("[ ! ] Finished Migration `%s` Took %s \n", migrator.Name, time.Since(startedMigrationAt).String())
	}

	if err := txn.Commit(); err != nil {
		return err
	}

	log.Printf("[ ! ] Finished! Took %s \n", time.Since(startedConverterAt).String())
	return nil
}
