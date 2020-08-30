package postgres

import (
	"database/sql"
	"fmt"
)

type Postgres struct {
	db *sql.DB
}

type Options struct {
	Host        string
	DBUser      string
	DBPass      string
	DBName      string
	DisableSSL  bool
	DatabaseURL string
}

func New(opts Options) (*Postgres, error) {
	if opts.DatabaseURL == "" {
		sslmode := "require"
		if opts.DisableSSL {
			sslmode = "disable"
		}
		opts.DatabaseURL = fmt.Sprintf("host=%s user=%s password=%s port=5432 dbname=%s sslmode=%s",
			opts.Host,
			opts.DBUser,
			opts.DBPass,
			opts.DBName,
			sslmode,
		)
	}
	db, err := sql.Open("postgres", opts.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}
