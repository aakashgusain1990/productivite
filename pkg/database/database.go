package database

import "github.com/uptrace/bun"

type Database interface {
	GetDB() *bun.DB
}