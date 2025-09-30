package app

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/wolf1848/taxiportal/api"
	"github.com/wolf1848/taxiportal/drivers"
	"github.com/wolf1848/taxiportal/logger"
	"github.com/wolf1848/taxiportal/model"
	"github.com/wolf1848/taxiportal/repository"
	"github.com/wolf1848/taxiportal/service"
)

func New() {
	cfg := newConfig()
	log := logger.New("web", cfg.Loglevel)

	db := drivers.NewPostgres(&cfg.Database)

	repositories := repository.NewRepositories(db)
	services := service.NewServices(cfg, repositories, log)

	server := api.New(services)

	go func() {
		server.Start(cfg.Server.Host + ":" + cfg.Server.Port)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error(err.Error())
	}

	if err := drivers.ShutdownPostgres(ctx, db); err != nil {
		log.Error(err.Error())
	}

}

func newConfig() *model.AppApiConfig {
	var env string
	var err error

	flag.StringVar(&env, "env", "production", "Runtime environment variable (default: production)")
	flag.Parse()

	if err = godotenv.Load("./config/" + env + "/.env"); err != nil {
		panic(err.Error())
	}

	v := viper.NewWithOptions(viper.KeyDelimiter("_"))

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config/" + env + "/")
	v.AutomaticEnv()

	if err = v.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	conf := &model.AppApiConfig{}

	if err = v.UnmarshalExact(conf); err != nil {
		panic(err.Error())
	}

	return conf
}
