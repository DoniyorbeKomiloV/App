package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageInterface, logger logger.LoggerI) {
	NewHandler := handler.NewHandler(cfg, storage, logger)

	r.POST("/books", NewHandler.CreateBook)
	r.GET("/books/:id", NewHandler.GetByIdBook)
	r.GET("/books", NewHandler.GetListBooks)
	r.PUT("/books", NewHandler.UpdateBook)
	r.DELETE("/books/:id", NewHandler.DeleteBook)

	r.POST("/users", NewHandler.CreateUser)
	r.GET("/users/:id", NewHandler.GetByIdUser)
	r.GET("/users", NewHandler.GetListUsers)
	r.PUT("/users", NewHandler.UpdateUser)
	r.DELETE("/users/:id", NewHandler.DeleteUser)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
