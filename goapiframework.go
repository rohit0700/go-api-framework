package goapiframework

import (
	"database/sql"
	"github.com/rohit0700/go-api-framework/config"
	"github.com/rohit0700/go-api-framework/database"
	"github.com/rohit0700/go-api-framework/errors"
	"github.com/rohit0700/go-api-framework/http"
	"github.com/rohit0700/go-api-framework/logging"
)

type Framework struct {
	Config *config.Config
	DB     *sql.DB
}

func NewFramework() (*Framework, error) {
	// Initialize configuration
	cfg, err := config.InitConfig()
	if err != nil {
		return nil, err
	}

	// Initialize logging
	// Note: You may customize logging settings based on your needs
	logging.InfoLogger.Println("Initializing framework...")

	// Initialize database
	db, err := database.InitDatabase(cfg)
	if err != nil {
		logging.ErrorLogger.Fatal("Error initializing database:", err)
		return nil, errors.NewError(errors.ErrDatabase)
	}

	// Initialize HTTP router
	router := http.InitRouter(cfg)

	logging.InfoLogger.Println("Framework initialized successfully", router)

	return &Framework{
		Config: cfg,
		DB:     db,
	}, nil
}
