package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/dd3v/snippets.ninja/internal/auth"
	"github.com/dd3v/snippets.ninja/internal/config"
	"github.com/dd3v/snippets.ninja/internal/errors"
	"github.com/dd3v/snippets.ninja/internal/rbac"
	"github.com/dd3v/snippets.ninja/internal/snippet"
	"github.com/dd3v/snippets.ninja/internal/user"
	"github.com/dd3v/snippets.ninja/pkg/accesslog"
	"github.com/dd3v/snippets.ninja/pkg/dbcontext"
	"github.com/dd3v/snippets.ninja/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
	"time"
)

var Version = "1.0 beta"

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "../../config/local.yml", "path to config file")
}

func main() {
	flag.Parse()
	logger := log.New([]string{
		"stdout",
	})
	cfg, err := config.Load(configPath)
	logger.Infof("Config path: %s", configPath)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		os.Exit(-1)
	}
	mysql, err := dbx.MustOpen("mysql", cfg.DatabaseDNS)
	if err != nil {
		logger.Info(cfg.DatabaseDNS)
		logger.Errorf("DB connection error %v", err)
	}
	mysql.ExecLogFunc = func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB query successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB query error: %v", err)
		}
	}
	mysql.QueryLogFunc = func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
		if err == nil {
			logger.With(ctx, "duration", t.Milliseconds(), "sql", sql).Info("DB query successful")
		} else {
			logger.With(ctx, "sql", sql).Errorf("DB query error: %v", err)
		}
	}
	defer func() {
		if err := mysql.Close(); err != nil {
			logger.Error(err)
		}
	}()
	db := dbcontext.New(mysql)
	rbac := rbac.New()

	jwtAuthMiddleware := auth.GetJWTMiddleware(cfg.JWTSigningKey)
	router := routing.New()
	router.Use(
		accesslog.Handler(logger),
		content.TypeNegotiator(content.JSON),
		errors.Handler(),
	)
	apiGroup := router.Group("/api")
	apiGroup.Get("/version", func(c *routing.Context) error {
		return c.Write(Version)
	})
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	user.NewHTTPHandler(apiGroup.Group("/v1"), jwtAuthMiddleware, userService)
	auth.NewHTTPHandler(apiGroup.Group("/v1"), jwtAuthMiddleware, auth.NewService(cfg.JWTSigningKey, auth.NewRepository(db), logger))
	snippet.NewHTTPHandler(apiGroup.Group("/v1"), jwtAuthMiddleware, snippet.NewService(
		snippet.NewRepository(db),
		rbac,
	))
	address := fmt.Sprintf(":%v", cfg.BindAddr)
	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("http server error: %s", err)
		os.Exit(-1)
	}
}
