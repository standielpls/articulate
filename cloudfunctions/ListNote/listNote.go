package p

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

var v variables

type variables struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
}

var pg *Postgres

func init() {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		panic("DB_USER not provided")
	}
	v.DBUser = dbUser
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		panic("DB_PASS not provided")
	}
	v.DBPass = dbPass
	dbHost := os.Getenv("DB_HOST")
	if dbUser == "" {
		panic("DB_HOST not provided")
	}
	v.DBHost = dbHost
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic("DB_NAME not provided")
	}
	v.DBName = dbName

	var once sync.Once
	once.Do(func() {
		p, err := connectPostgres(v)
		if err != nil {
			panic(fmt.Sprintf("unable to connect to Postgres: %s", err.Error()))
		}
		if err := p.Ping(); err != nil {
			panic(fmt.Sprintf("unable to ping db: %s", err.Error()))
		}
		pg = p
	})
}

type Request struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Response struct {
	Articles []Article `json:"articles"`
	Message  string    `json:"message"`
	Ok       bool      `json:"ok"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	ctx := context.Background()
	articles, err := pg.ListByUserID(ctx, r.URL.Query().Get("user_id"))
	if err != nil {
		data := Response{
			Message: err.Error(),
			Ok:      false,
		}
		b, _ := json.Marshal(data)
		w.Write(b)
		w.WriteHeader(400)
	}

	data := Response{
		Articles: articles,
		Ok:       true,
	}
	b, _ := json.Marshal(data)
	w.Write(b)
	w.WriteHeader(200)
}

type Postgres struct {
	db *sql.DB
}

func connectPostgres(v variables) (*Postgres, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s", v.DBHost, v.DBUser, v.DBPass, v.DBName))
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}

type Article struct {
	URL       string    `json:"url"`
	Article   string    `json:"article"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Articler interface {
	ListByUserID(ctx context.Context, userid string) ([]Article, error)
}

func (p *Postgres) Ping() error {
	return p.db.Ping()
}
func (p *Postgres) ListByUserID(ctx context.Context, id string) ([]Article, error) {
	rows, err := p.db.QueryContext(ctx, `
		SELECT url, article, comment, created_at, updated_at
		FROM article
		WHERE user_id=$1
	`, id)
	if err != nil {
		return nil, err
	}

	var as []Article
	for rows.Next() {
		var a Article
		err := rows.Scan(&a.URL, &a.Article, &a.Comment, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}

		as = append(as, a)
	}

	return as, nil
}
