package server

import (
	memberHttp "social_media/internal/member/delivery/http"
	memberRepo "social_media/internal/member/repository"
	memberUsecase "social_media/internal/member/usecase"
	"social_media/internal/middleware"
	productHttp "social_media/internal/product/delivery/http"
	productRepo "social_media/internal/product/repository"
	productUsecase "social_media/internal/product/usecase"

	docs "social_media/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	docs.SwaggerInfo.Title = s.cfg.Swagger.Title
	docs.SwaggerInfo.Description = s.cfg.Swagger.Description
	docs.SwaggerInfo.Version = s.cfg.Swagger.Version
	docs.SwaggerInfo.BasePath = s.cfg.Swagger.BasePath
	docs.SwaggerInfo.Host = s.cfg.Swagger.Host

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	mw := middleware.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger)
	apiGroup := e.Group("/api/v1", mw.RequestLoggerMiddleware)

	memberGroup := apiGroup.Group("/members")
	productsGroup := apiGroup.Group("/products")

	memberRepo := memberRepo.NewMemberRepository(s.db)
	memberUC := memberUsecase.NewMemberUsecase(memberRepo)

	productRepo := productRepo.NewMySQLProductRepository(s.db)
	productUC := productUsecase.NewProductUsecase(productRepo)

	memberHttp.MapMemberRoute(memberGroup, s.logger, memberUC)
	productHttp.MapProductRoutes(productsGroup, s.logger, productUC)
	return nil
}
