package converter

import (
	"context"
	"fmt"

	"github.com/tankbusta/gocrackdb-converter/lib/ent"
	"github.com/tankbusta/gocrackdb-converter/lib/ent/user"
	"github.com/tankbusta/gocrackdb-converter/lib/oldschema"

	"github.com/asdine/storm"
	"github.com/google/uuid"
)

type boltUser struct {
	ID             int64 `storm:"id,increment"`
	DocVersion     float32
	oldschema.User `storm:"inline"`
}

func convertBaseUsers(ctx context.Context, from *storm.DB, to *ent.Tx) error {
	var records []*ent.UserCreate

	if err := from.From("users").Select().Each(new(boltUser), func(record interface{}) error {
		u := record.(*boltUser)
		uid, err := uuid.Parse(u.UserUUID)
		if err != nil {
			return fmt.Errorf("user `%s` has a bad id: %w", u.Username, err)
		}

		var enabled bool
		if u.Enabled != nil && !*u.Enabled {
			enabled = false
		}

		records = append(
			records,
			to.User.Create().
				SetID(uid).
				SetUsername(u.Username).
				SetEnabled(enabled).
				SetEmailAddress(u.EmailAddress).
				SetIsSuperUser(u.IsSuperUser).
				SetCreatedAt(u.CreatedAt),
		)
		return nil
	}); err != nil {
		return err
	}

	return to.User.CreateBulk(records...).
		OnConflictColumns(user.FieldID).
		DoNothing().
		Exec(ctx)
}
