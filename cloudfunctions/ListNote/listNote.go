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

	"github.com/gomodule/redigo/redis"
	_ "github.com/lib/pq"
)

var v variables

type variables struct {
	DBUser    string
	DBPass    string
	DBHost    string
	DBName    string
	RedisAddr string
	RedisPass string
}

var pg *Postgres

func ok() {
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
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		panic("REDIS_ADDR not provided")
	}
	v.RedisAddr = redisAddr
	redisPass := os.Getenv("REDIS_PASS")
	if redisPass == "" {
		panic("REDIS_PASS not provided")
	}
	v.RedisPass = redisPass

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
		ORDER BY updated_at DESC
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

type Redis struct {
	con redis.Conn
}

func NewRedis(addr, pw string) (*Redis, error) {
	conn, err := redis.Dial("tcp", addr+":6379", redis.DialPassword(pw))
	if err != nil {
		return nil, fmt.Errorf("unable to dial: %s", err.Error())
	}
	return &Redis{conn}, nil
}
func (r *Redis) GetNote(ctx context.Context, key string) ([]Article, error) {
	res, err := redis.Bytes(r.con.Do("LRANGE", key, 0, -1))
	if err != nil {
		return nil, err
	}

	var a []Article
	err = json.Unmarshal(res, &a)
	return a, err
}

func (r *Redis) CreateNote(ctx context.Context, id string, a Article) error {

	res, err := redis.Bytes(r.con.Do("SET", id, a))
	if err != nil {
		return nil, err
	}

	var a []Article
	err = json.Unmarshal(res, &a)
	return a, err
}
