package main

import (
	"database/sql"
	"fmt"
	"main/internal/pkg/config"
	"main/internal/pkg/user"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const AppName = "animal-chipization"

func main() {
	cfg, err := config.Load("./config/config.yml")
	if err != nil {
		panic(err)
	}
	pgDB, err := openPgDB(cfg.Postgres)
	if err != nil {
		panic(err)
	}
	defer pgDB.Close()

	app := fiber.New(fiber.Config{
		AppName: AppName,
	})
	var userHandler user.Handler
	{
		rep := user.NewPostgresRepo(pgDB)
		svc := user.NewService(rep)
		userHandler = user.NewHandler(svc)
		userHandler.Register(app)
	}

	_ = app.Listen(":8080")
}

func openPgDB(cfg config.Postgres) (*sql.DB, error) {
	pgDB, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Database,
			cfg.SSLMode,
		))
	if err != nil {
		return nil, err
	}
	err = pgDB.Ping()
	if err != nil {
		pgDB.Close()
		return nil, err
	}

	return pgDB, nil
}
