package service

import (
	"go.uber.org/zap"
	"log"
	"os"
)

type LoggerService struct {
	Log *zap.Logger
}

func NewLoggerService() *LoggerService {
	logger, err := zap.NewProduction()
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1] == "debug"{
		logger, err = zap.NewDevelopment()
		if nil != err {
			log.Println(err)
			os.Exit(1)
		}
	}

	defer logger.Sync()

	logger.Debug("Logger Service Initialized")
	return &LoggerService{
		Log: logger,
	}
}