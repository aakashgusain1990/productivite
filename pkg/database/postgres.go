package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/aakashgusain1990/productivite/pkg/entities"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type PGDatabase struct {
	DB *bun.DB
}

var (
	once sync.Once
	dbInstance *PGDatabase
)

func NewPGDatabase(conf *entities.Config) *PGDatabase {
	once.Do(func() {
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
				conf.Env.Local.Database.User,
				conf.Env.Local.Database.Password,
				conf.Env.Local.Database.Host,
				conf.Env.Local.Database.Port,
				conf.Env.Local.Database.DB)
		config, err := pgx.ParseConfig(dsn)
		if err != nil {
			panic(err)
		}
		pgDB := stdlib.OpenDB(*config)
		fmt.Printf("%+v", pgDB)
		dbInstance = &PGDatabase{DB: bun.NewDB(pgDB, pgdialect.New())}
		fmt.Printf("%+v", dbInstance)
		ctx := context.Background()
		res, err := dbInstance.DB.ExecContext(ctx, "SELECT * FROM public.expenses;")

		fmt.Print(res)
	})
	return dbInstance
}

func (pgDatabase *PGDatabase) GetDB() *bun.DB {
	return dbInstance.DB
}