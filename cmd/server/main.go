package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"

	"github.com/dd3v/snippets.page.backend/internal/auth"
	"github.com/dd3v/snippets.page.backend/internal/config"
	"github.com/dd3v/snippets.page.backend/internal/errors"
	"github.com/dd3v/snippets.page.backend/internal/rbac"
	"github.com/dd3v/snippets.page.backend/internal/snippet"
	"github.com/dd3v/snippets.page.backend/internal/user"
	"github.com/dd3v/snippets.page.backend/pkg/dbcontext"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "../../config/app.toml", "path to config file")
}

func main() {
	config := config.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	mysql, err := dbx.MustOpen("mysql", config.DatabaseDNS)
	if err != nil {
		fmt.Printf("mysql connection error: %s", err)
	}
	defer func() {
		if err := mysql.Close(); err != nil {
			fmt.Printf("mysql runtime error: %s", err)
		}
	}()
	db := dbcontext.New(mysql)
	rbac := rbac.New()
	jwtAuthHandler := auth.Handler(config.JWTSigningKey)
	router := routing.New()
	router.Use(
		content.TypeNegotiator(content.JSON),
		errors.Handler(),
	)
	apiGroup := router.Group("/api")
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	user.NewHTTPHandler(apiGroup.Group("/v1"), jwtAuthHandler, userService)
	auth.NewHTTPHandler(apiGroup.Group("/v1"), jwtAuthHandler, auth.NewService(config.JWTSigningKey, auth.NewRepository(db)))
	snippet.NewHTTPHandler(apiGroup.Group("/v1"), jwtAuthHandler, snippet.NewService(
		snippet.NewRepository(db),
		rbac,
	))
	address := fmt.Sprintf(":%v", config.BindAddr)
	httpServer := &http.Server{
		Addr:    address,
		Handler: router,
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("http server error: %s", err)
		os.Exit(-1)
	}
}
