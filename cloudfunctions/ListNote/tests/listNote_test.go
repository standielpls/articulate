package test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	p "github.com/standielpls/articulate/cloudfunctions/ListNote"
)

func TestGetNote(t *testing.T) {
	r, err := p.NewRedis("35.221.31.170", "refined-tarn-willow-mod-unwonted")

	ctx := context.Background()

	notes, err := r.GetNote(ctx, "hello")
	if err != nil {
		t.Fatalf("unable to get note: %s", err.Error())
	}

	expRes := []p.Article{
		{
			URL:     "abc.com",
			Article: "Hello World",
			Comment: "Cool!",
		},
	}
	if !cmp.Equal(notes, expRes) {
		t.Fatalf("unexpected result returned: %s", cmp.Diff(notes, expRes))
	}
}
