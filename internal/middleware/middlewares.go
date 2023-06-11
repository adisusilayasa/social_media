package middleware

import (
	"social_media/config"
	"social_media/pkg/zap"
)

type MiddlewareManager struct {
	cfg     *config.Config
	origins []string
	logger  zap.Logger
}

func NewMiddlewareManager(cfg *config.Config, origins []string, logger zap.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg: cfg, origins: origins, logger: logger}
}
