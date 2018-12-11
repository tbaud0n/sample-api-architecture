package main

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/tbaud0n/sample-api-architecture/pkg/config"
	api_http "github.com/tbaud0n/sample-api-architecture/pkg/http"
	"github.com/tbaud0n/sample-api-architecture/pkg/logger"
	"github.com/tbaud0n/sample-api-architecture/pkg/storage/postgresql"
)

var (
	configFile  = `config.yaml`
	logFilepath = `var/log/api.log`
)

// Route describe a route
type Route struct {
	Name    string
	Path    string
	Method  string
	Handler http.Handler
}

func main() {

	config, err := config.NewConfig(configFile)
	if err != nil {
		logger.LogFatal(err)
	}

	initLogger()

	db, err := sql.Open("postgres", config.DatabaseDSN)
	if err != nil {
		logger.LogFatal(err)
	}
	defer db.Close()

	us := &postgresql.UserService{DB: db}

	routes := []Route{
		{
			Name:   `userQuery`,
			Method: http.MethodGet,
			Handler: api_http.UserQueryHandler{
				UserStorageService:         us,
				QueryFilterFromHTTPRequest: postgresql.QueryFilterFromHTTPRequest,
			},
			Path: `/users`,
		},
		{
			Name:    `userGet`,
			Method:  http.MethodGet,
			Handler: api_http.UserGetHandler{UserStorageService: us},
			Path:    `/users/{id:[0-9]+}`,
		},
	}

	router := mux.NewRouter()

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	logger.LogInfo("HTTP server running. Listening on http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", router)
}

func initLogger() {
	lf := getLogFile()
	mw := io.MultiWriter(os.Stdout, lf)
	logger.SetOutput(mw)
}

func getLogFile() (lf *os.File) {
	err := os.MkdirAll(filepath.Dir(logFilepath), 0774)
	if err != nil {
		err = errors.Wrap(err, "Error creating log file parent directory : ")
		logger.LogFatal(err)
	}

	lf, err = os.OpenFile(logFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0774)
	if err != nil {
		err = errors.Wrap(err, "Error opening log file")
		logger.LogFatal(err)
	}

	return
}
