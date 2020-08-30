package postgres

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ryanfowler/uuid"
	art "github.com/standielpls/articulate/server"

	_ "github.com/lib/pq"
)

func connectPg() *Postgres {
	pg, err := New(Options{
		Host:       "localhost",
		DBUser:     "postgres",
		DBName:     "articulate_test",
		DisableSSL: true,
	})
	if err != nil {
		panic(err)
	}
	return pg
}

func cleanTable(t *testing.T, db *sql.DB) {
	t.Helper()

	_, err := db.Exec(`
		DELETE FROM articles;
		DELETE FROM users;
	`)
	if err != nil {
		panic(err)
	}
}

func TestCreateGetNote(t *testing.T) {
	p := connectPg()
	cleanTable(t, p.db)

	ctx := context.Background()

	id, _ := uuid.NewV4()
	note := art.Article{
		ID:      id.String(),
		URL:     "url.com",
		Article: "don't be evil",
		Comment: "evil is not nice!",
	}

	userID, err := p.CreateUser(ctx, "1")
	if err != nil {
		t.Fatalf("unable to create user: %s", err.Error())
	}

	note.UserID = userID

	noteID, err := p.CreateNote(ctx, note)
	if err != nil {
		t.Fatalf("unable to create note: %s", err.Error())
	}
	note.ID = noteID

	notes, err := p.ListByUserID(ctx, note.UserID)
	if err != nil {
		t.Fatalf("unable to list notes: %s", err.Error())
	}

	if len(notes) < 1 {
		t.Fatalf("unexpected number of notes returned: %s", cmp.Diff(len(notes), 1))
	}

	opts := cmpopts.IgnoreFields(art.Article{}, "CreatedAt", "UpdatedAt")
	diff := cmp.Diff(notes[0], note, opts)
	if diff != "" {
		t.Fatalf("unexpected note returned: %s", diff)
	}
}
