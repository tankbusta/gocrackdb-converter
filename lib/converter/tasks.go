package converter

import (
	"context"
	"fmt"
	"log"

	"github.com/tankbusta/gocrackdb-converter/lib/ent"
	"github.com/tankbusta/gocrackdb-converter/lib/ent/task"
	"github.com/tankbusta/gocrackdb-converter/lib/oldschema"

	"github.com/asdine/storm"
	"github.com/google/uuid"
)

type boltCrackTask struct {
	DocVersion     float32
	oldschema.Task `storm:"inline"`
}

func convertBaseTasks(ctx context.Context, from *storm.DB, to *ent.Tx) error {
	return from.From("tasks").Select().Each(new(boltCrackTask), func(record interface{}) error {
		t := record.(*boltCrackTask)
		tid, err := uuid.Parse(t.TaskID)
		if err != nil {
			return fmt.Errorf("task `%s` has a bad user id: %w", t.TaskID, err)
		}

		uid, err := uuid.Parse(t.CreatedByUUID)
		if err != nil {
			return fmt.Errorf("task `%s` has a bad user id: %w", t.TaskID, err)
		}

		q := to.Task.Create().
			SetID(tid).
			SetName(t.TaskName).
			SetStatus(t.Status).
			SetCreatedByID(uid).
			SetCreatedAt(t.CreatedAt).
			SetLastUpdatedAt(t.LastUpdatedAt).
			SetAssignedToHost(t.AssignedToHost).
			SetNillableComment(t.Comment).
			SetNillableCaseCode(t.CaseCode).
			SetNumberCracked(t.NumberCracked).
			SetNumberPasswords(t.NumberPasswords).
			SetNillableError(t.Error)

		switch v := t.EnginePayload.(type) {
		case nil:
			// Do nothing
		case map[string]interface{}:
			q = q.SetEnginePayload(v)
		default:
			log.Printf("[ ? ] Unknown Type for Engine Payload: %T\n", v)
		}

		return q.OnConflictColumns(task.FieldID).
			DoNothing().
			Exec(ctx)
	})
}
