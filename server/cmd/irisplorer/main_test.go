package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/irisnet/irisplorer.io/server/modules/store"
	"github.com/irisnet/irisplorer.io/server/modules/tools"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain2(t *testing.T) {
	//viper.Set("node", "tcp://localhost:46657")
	//viper.Set("node", "tcp://47.104.155.125:46757")
	viper.Set("node", "tcp://116.62.62.39:11657")
	viper.Set(tools.InitConnectionNum, 100)
	viper.Set(tools.MaxConnectionNum, 200)
	viper.Set(tools.ChainId, "pangu")
	viper.Set(tools.SyncCron, "@every 3s")
	viper.Set(store.MgoUrl, "localhost:27017")
	tools.Init()
	StartWatch(nil, nil)

	time.Sleep(1 * time.Hour)
}

func TestRunRestServer(t *testing.T) {
	viper.Set("node", "tcp://47.104.155.125:46757")
	viper.Set(tools.InitConnectionNum, 5)
	viper.Set(tools.MaxConnectionNum, 10)

	tools.Init()
	//store.Init("localhost:27017")

	router := mux.NewRouter()
	// latest
	AddRoutes(router)
	http.ListenAndServe(":8998",
		handlers.LoggingHandler(os.Stdout, handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}))(router)))

	time.Sleep(1 * time.Hour)
}
