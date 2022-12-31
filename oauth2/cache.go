package oauth2

import (
	"github.com/disism/oauth2/config"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"sync"
	"time"
)

type Cache struct {
	Ctx    context.Context
	Client *redis.Client
	Err    error
}

func NewCache(db int) *Cache {
	if err := NewRdb().Dial(db); err != nil {
		return &Cache{Err: err}
	}
	return &Cache{Ctx: context.Background(), Client: rdb, Err: nil}
}

var (
	once sync.Once
	rdb  *redis.Client
)

type Redis interface {
	Dial(db int) error
}

type option struct {
	addr     string
	password string
}

func NewRdb() *option {
	return &option{
		addr:     config.GetRedisAddr(),
		password: config.GetRedisPassword(),
	}
}

func (r *option) Dial(db int) error {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Network:            "tcp",
			Addr:               r.addr,
			Dialer:             nil,
			OnConnect:          nil,
			Username:           "",
			Password:           r.password,
			DB:                 db,
			MaxRetries:         0,
			MinRetryBackoff:    0,
			MaxRetryBackoff:    0,
			DialTimeout:        5 * time.Second,
			ReadTimeout:        3 * time.Second,
			WriteTimeout:       3 * time.Second,
			PoolSize:           15,
			MinIdleConns:       10,
			MaxConnAge:         0,
			PoolTimeout:        4 * time.Second,
			IdleTimeout:        0,
			IdleCheckFrequency: 0,
			TLSConfig:          nil,
			Limiter:            nil,
		})
	})
	return nil
}

type AuthCode interface {
	SETCODE(code, userId string) error
	GETUSERID(code string) string
}

func (r *Cache) SETCODE(code, userId string) error {
	if err := r.Client.Set(r.Ctx, code, userId, 1200*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func (r *Cache) GETUSERID(code string) string {
	return r.Client.Get(r.Ctx, code).Val()
}
