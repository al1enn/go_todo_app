package main

import (
	"os"

	todo "github.com/al1enn/go_todo_app"
	"github.com/al1enn/go_todo_app/pkg/handler"
	"github.com/al1enn/go_todo_app/pkg/repository"
	"github.com/sirupsen/logrus"

	"github.com/al1enn/go_todo_app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	// logrus.SetFormatter(new(logrus.JSONFormatter)) // json Format For logs
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error occured while reading config: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occured while loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error occured while connecting to db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	server := new(todo.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil { //deFault port oF http server  is 80
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
