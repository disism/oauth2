package oauth2

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db *gorm.DB
)

type IDatabase interface {
	Dial() error
}

type Database struct {
	addr string
}

func NewDatabase() *Database {
	var (
		host     = viper.GetString("db.host")
		port     = viper.GetString("db.port")
		user     = viper.GetString("db.user")
		password = viper.GetString("db.password")
		dbName   = viper.GetString("db.database")
		sslMode  = viper.GetString("db.ssl")
		timeZone = viper.GetString("db.timezone")
	)

	addr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbName,
		port,
		sslMode,
		timeZone,
	)

	return &Database{addr: addr}
}

func (d *Database) Dial() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("CONNECT_TO_COCKROACH_DB_TIMEOUT")
		default:
			dbs, err := gorm.Open(postgres.Open(d.addr), &gorm.Config{})
			if err != nil {
				return err
			}
			log.Println("DATABASE_CONNECTION_SUCCESSFUL")
			db = dbs
			return nil
		}
	}
}

func InitDatabase() error {
	if err := NewDatabase().Dial(); err != nil {
		return err
	}
	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
