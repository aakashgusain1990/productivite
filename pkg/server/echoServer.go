package server

import (
	"fmt"

	"github.com/aakashgusain1990/productivite/pkg/entities"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/uptrace/bun"
)

type echoServer struct {
	server *echo.Echo
	db *bun.DB
	conf *entities.Config
}

func NewServer(conf *entities.Config, pgDatabase *bun.DB) *echoServer{
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	return &echoServer{
		server: e,
		db: pgDatabase,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	addr := fmt.Sprintf(":%d", s.conf.Env.Local.Port)
	s.server.Logger.Fatal(s.server.Start(addr))
}