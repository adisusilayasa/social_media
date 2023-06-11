package main

import (
	"context"
	"log"
	"social_media/config"
	"social_media/config/db"
	"social_media/internal/server"
	"social_media/pkg/zap"
	"time"
)

const (
	ctxTimeout = 10
)

func main() {
	log.Println("Starting app server")

	cfgFile, err := config.LoadConfig("./config/config")
	if err != nil {
		log.Fatalf("Load config: ", err.Error())
	}

	cfg, err := config.ParseConfigDefault(cfgFile)
	if err != nil {
		log.Fatalf("Parse config: ", err.Error())
	}

	logger := zap.NewAppLogger(cfg)
	logger.InitLogger()
	logger.Infof("App Version: 1.0.0")

	_, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()

	db, err := db.InitDatabase(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	s := server.NewServer(cfg, logger, db)
	if err := s.Run(); err != nil {
		logger.Fatal(err.Error())
	}
}
