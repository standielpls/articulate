package postgres

import (
	"context"

	art "github.com/standielpls/articulate/server"
)

func (p *Postgres) CreateNote(ctx context.Context, note art.Article) (string, error) {
	var id string
	err := p.db.QueryRowContext(ctx, `
		INSERT INTO articles
		(id, url, comment, article, user_id)
		VALUES
		($1, $2, $3, $4, $5)
		RETURNING id
	`, note.ID, note.URL, note.Comment, note.Article, note.UserID).Scan(&id)
	return id, err
}

func (p *Postgres) ListByUserID(ctx context.Context, id string) ([]art.Article, error) {
	rows, err := p.db.QueryContext(ctx, `
		SELECT id, url, article, comment, user_id, created, updated
		FROM articles
		WHERE user_id=$1
		ORDER BY updated DESC
	`, id)
	if err != nil {
		return nil, err
	}

	var as []art.Article
	for rows.Next() {
		var a art.Article
		err := rows.Scan(&a.ID, &a.URL, &a.Article, &a.Comment, &a.UserID, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}
