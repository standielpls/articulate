package p

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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
	ID      string `json:"id"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var n Article
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		data := Response{
			Message: err.Error(),
			Ok:      false,
		}
		b, _ := json.Marshal(data)
		w.Write(b)
		w.WriteHeader(400)
		return
	}
	if err := n.isValid(); err != nil {
		data := Response{
			Message: err.Error(),
			Ok:      false,
		}
		b, _ := json.Marshal(data)
		w.Write(b)
		w.WriteHeader(400)
		return
	}

	id, err := pg.CreateNote(ctx, n)
	if err != nil {
		data := Response{
			Message: err.Error(),
			Ok:      false,
		}
		b, _ := json.Marshal(data)
		w.Write(b)
		w.WriteHeader(400)
		return
	}

	data := Response{
		ID: id,
		Ok: true,
	}
	b, _ := json.Marshal(data)
	w.Write(b)
	w.WriteHeader(200)
}

func (a *Article) isValid() error {
	if a.URL == "" {
		return errors.New("url must be provided")
	}
	if a.Article == "" {
		return errors.New("article must be provided")
	}

	if a.Comment == "" {
		return errors.New("comment must be provided")
	}

	if a.UserID == "" {
		return errors.New("user_id must be provided")
	}

	return nil
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
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Articler interface {
	ListByUserID(ctx context.Context, userid string) ([]Article, error)
}

func (p *Postgres) Ping() error {
	return p.db.Ping()
}
func (p *Postgres) CreateNote(ctx context.Context, note Article) (string, error) {
	var id string
	err := p.db.QueryRowContext(ctx, `
		INSERT INTO article
		(url, comment, article, user_id)
		VALUES
		($1, $2, $3, $4)
		RETURNING id
	`, note.URL, note.Comment, note.Article, note.UserID).Scan(&id)
	return id, err
}
