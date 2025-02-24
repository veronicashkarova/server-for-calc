package application

import (
	"fmt"
	"net/http"
	"os"
) 
type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

type Request struct {
	Expression string `json:"expression"`
}

type TaskRequest struct {
	ID int `json:"id"`
	Result int `json:"result"`
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", NewExpressionHandler)
	http.HandleFunc("/api/v1/expressions", ExpressionsHandler)
	http.HandleFunc("/api/v1/expressions/", IdHandler)
	http.HandleFunc("/internal/task", TaskHandler)
	fmt.Println("Server started")
	return http.ListenAndServe("", nil)
}
