package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Hyp9r/sequncer/domain/sequence"
	sequenceInfra "github.com/Hyp9r/sequncer/infrastructure/sequence"
	sequenceTransport "github.com/Hyp9r/sequncer/transport/sequence"
	"github.com/Netflix/go-env"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type AppConfig struct {
	Port        string `env:"PORT" envDefault:"8080"`
	SQLUsername string `env:"SQL_USERNAME"`
	SQLPassword string `env:"SQL_PASSWORD"`
	SQLDatabase string `env:"SQL_DATABASE"`
	SQLHost     string `env:"SQL_HOST" envDefault:"localhost"`
}

func main() {
	logger := zerolog.New(os.Stderr)
	fmt.Println("Hello, Webserver!")

	var Cfg AppConfig
	_, err := env.UnmarshalFromEnviron(&Cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to parse environment variables")
	}

	dbConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", Cfg.SQLUsername, Cfg.SQLPassword, Cfg.SQLDatabase, Cfg.SQLHost)
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		logger.Fatal().Err(err).Msg("error connecting to postgres")
	}
	defer db.Close()

	// initialize concrete db implementations
	sequenceRepo := sequenceInfra.NewSequenceRepository(db, &logger)

	// initalize service
	sequenceService := sequence.NewSequenceService(sequenceRepo)

	router := http.NewServeMux()
	sequenceTransport.NewController(router, sequenceService)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", Cfg.Port),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
	}
}
