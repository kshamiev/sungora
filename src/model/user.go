package model

import (
	"context"

	"sungora/lib/app"
	"sungora/lib/storage"
	"sungora/lib/storage/stpg"
	"sungora/lib/uuid"
	"sungora/src/typ"
)

// бизнес модель
type User struct {
	storage.Face
}

// NewUser
func NewUser() *User { return &User{&stpg.Storage{}} }

// Load
func (u *User) Load(ctx context.Context, id uuid.UUID) (*typ.Users, error) {
	s := app.NewSpan(ctx)
	s.StringAttribute("param1", "fantik")
	s.Int64Attribute("param2", 34)
	s.Float64Attribute("param3", 45.76)
	s.BoolAttribute("param4", true)
	defer s.End()

	us := &typ.Users{}

	if err := u.Gist().GetContext(ctx, us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	if err := u.Query(ctx).Get(us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	if err := u.QueryTx(ctx, func(qu storage.QueryTxEr) error {
		if err := qu.Get(us, "SELECT * FROM users WHERE id = $1", id); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return us, nil
}
