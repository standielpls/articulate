package postgres

import "context"

func (p *Postgres) CreateUser(ctx context.Context, userID string) (string, error) {
	var id string
	err := p.db.QueryRowContext(ctx, `
		INSERT INTO users (id)
		VALUES ($1)
		RETURNING id;
	`, userID).Scan(&id)
	return id, err
}
