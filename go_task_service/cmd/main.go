package main

import (
	"flag"
	"go-microservices-template/go_task_service/config"
	"go-microservices-template/pkg/logger"
	"log"
)

func main() {
	flag.Parse()

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName("WriterService")
}
