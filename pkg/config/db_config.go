package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

const (
	dbDriver = "mysql"
)

type DBConfig struct {
	// Host is
	Host string `envconfig:"DB_HOST" default:"db"`

	// Port is
	Port int16 `envconfig:"DB_PORT" default:"3306"`

	// DatabaseName is
	DBName string `envconfig:"DB_NAME" default:"isuumo"`

	// User is
	User string `envconfig:"DB_USER" default:"44smkn"`

	// Password is
	Password string `envconfig:"DB_PASSWORD" default:"nkms44"`
}

func (cfg *DBConfig) GetDB(logger *zap.Logger) (*sql.DB, error) {
	connStrTemplate := "%v:%v@tcp(%v:%v)/%v" // mysql
	connStr := fmt.Sprintf(connStrTemplate, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		return nil, err
	}

	for {
		err = db.Ping()
		if err == nil {
			logger.Info("DB connection is OK", zap.String("dbhost", cfg.Host))
			break
		}
		time.Sleep(3 * time.Second)
		continue
	}

	// Parameter Set
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5)

	return db, nil
}
