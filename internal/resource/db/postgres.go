package db

import (
	"errors"
	"fmt"
	"rc-tournament/cmd/conf"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres Databse struct to use for interactions with Postgresql
type Postgres struct {
	db *gorm.DB

	conf *conf.Postgresql
}

// NewPostgres Factory Postgres struct
func NewPostgres(c *conf.Postgresql) *Postgres {
	return &Postgres{conf: c}
}

// Open Opens a new connection to Postgres with predefined configuration
func (p *Postgres) Open() error {
	var (
		e   error
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s",
			p.conf.Host,
			p.conf.User,
			p.conf.Password,
			p.conf.DBName,
			p.conf.Port,
		)
	)

	p.db, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if e != nil {
		return e
	}

	d, _ := p.db.DB()

	d.SetMaxIdleConns(1)
	d.SetMaxOpenConns(100)
	d.SetConnMaxIdleTime(4 * time.Minute)
	d.SetConnMaxLifetime(60 * time.Second)

	return nil
}

// Close closes connection to Postgres
func (p *Postgres) Close() error {
	if p.db == nil {
		return errors.New("db is nil")
	}

	d, _ := p.db.DB()

	return d.Close()
}

// Create creates a model on DB
func (p *Postgres) Create(model interface{}) (interface{}, error) {
	if tx := p.db.Create(model); tx.Error != nil {
		return model, tx.Error
	}
	return model, nil
}
