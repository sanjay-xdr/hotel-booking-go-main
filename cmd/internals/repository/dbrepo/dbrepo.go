package dbrepo

import (
	"database/sql"

	"github.com/sanjay-xdr/cmd/internals/config"
	"github.com/sanjay-xdr/cmd/internals/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
