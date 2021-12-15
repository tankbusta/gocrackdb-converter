package converter

import (
	"context"
	"fmt"
	"log"

	"github.com/tankbusta/gocrackdb-converter/lib/ent"
	"github.com/tankbusta/gocrackdb-converter/lib/oldschema"

	"github.com/asdine/storm"
	"github.com/schollz/progressbar/v3"
)

type boltCrackedHash struct {
	ID                    int64 `storm:"id,increment"`
	DocVersion            float32
	oldschema.CrackedHash `storm:"inline"`
}

// MaxNumberPerInsert is the maximum number of records that can be inserted in a batch.
// See comment below. There's no scientific reason why this number is 64 it worked well
// for our instance of GoCrack that had close to 7m records
const MaxNumberPerInsert = 64

func convertPasswords(ctx context.Context, from *storm.DB, to *ent.Tx) error {
	// Fetch the number of tasks we just indexed
	count, err := to.Task.Query().Count(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch # of tasks: %w", err)
	}

	log.Printf("[ ! ] Preparing to pull passwords from %d tasks\n", count)
	ids, err := to.Task.Query().Unique(true).IDs(ctx)
	if err != nil {
		return fmt.Errorf("failed to fetch task ids: %w", err)
	}

	bar := progressbar.Default(int64(count), "Cracked Hash Progress")
	for _, id := range ids {
		var records []*ent.CrackedPasswordCreate
		if err := from.From("tasks", id.String(), "results").Select().Each(new(boltCrackedHash), func(record interface{}) error {
			u := record.(*boltCrackedHash)

			records = append(
				records,
				to.CrackedPassword.Create().
					SetRelatedTaskID(id).
					SetHash(u.Hash).
					SetValue(u.Value).
					SetCrackedAt(u.CrackedAt),
			)

			return nil
		}); err != nil {
			return fmt.Errorf("failed to fetch cracked passwords for task %s: %w", id.String(), err)
		}

		// Iterate over the records and insert them in batches to avoid
		// errors like "parameters but PostgreSQL only supports 65535 parameters"
		for i := 0; i < len(records); i += MaxNumberPerInsert {
			j := i + MaxNumberPerInsert
			if j > len(records) {
				j = len(records)
			}

			// Process the batch.
			if err := to.CrackedPassword.CreateBulk(records[i:j]...).Exec(ctx); err != nil {
				return err
			}
		}

		bar.Add(1)
	}

	bar.Close()

	return nil
}
