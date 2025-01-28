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

	r.Use(customCORSMiddleware())
	r.Use(MaxAllowed(1000))

	r.POST("/books", NewHandler.Validate, NewHandler.CreateBook)
	r.GET("/books/:id", NewHandler.Validate, NewHandler.GetByIdBook)
	r.GET("/books", NewHandler.Validate, NewHandler.GetListBooks)
	r.PUT("/books", NewHandler.Validate, NewHandler.UpdateBook)
	r.DELETE("/books/:id", NewHandler.Validate, NewHandler.DeleteBook)

	r.POST("/users", NewHandler.Validate, NewHandler.CreateUser)
	r.GET("/users/:id", NewHandler.Validate, NewHandler.GetByIdUser)
	r.GET("/users", NewHandler.Validate, NewHandler.GetListUsers)
	r.PUT("/users", NewHandler.Validate, NewHandler.UpdateUser)
	r.DELETE("/users/:id", NewHandler.Validate, NewHandler.DeleteUser)

	r.POST("/orders", NewHandler.Validate, NewHandler.CreateOrder)
	r.GET("/orders/:id", NewHandler.Validate, NewHandler.GetByIdOrder)
	r.GET("/orders", NewHandler.Validate, NewHandler.GetListOrders)
	r.PUT("/orders", NewHandler.Validate, NewHandler.UpdateOrder)
	r.DELETE("/orders/:id", NewHandler.Validate, NewHandler.DeleteOrder)

	r.POST("/order_items", NewHandler.Validate, NewHandler.CreateOrderItem)
	r.GET("/order_items/:id", NewHandler.Validate, NewHandler.GetByIdOrderItem)
	r.GET("/order_items", NewHandler.Validate, NewHandler.GetListOrderItems)
	r.PUT("/order_items", NewHandler.Validate, NewHandler.UpdateOrderItem)
	r.DELETE("/order_items/:id", NewHandler.Validate, NewHandler.DeleteOrderItem)

	r.POST("/categories", NewHandler.Validate, NewHandler.CreateCategory)
	r.GET("/categories/:id", NewHandler.Validate, NewHandler.GetByIdCategory)
	r.GET("/categories", NewHandler.Validate, NewHandler.GetListCategories)
	r.PUT("/categories", NewHandler.Validate, NewHandler.UpdateCategory)
	r.DELETE("/categories/:id", NewHandler.Validate, NewHandler.DeleteCategory)

	r.POST("/upload", NewHandler.HandleUpload)

	r.POST("/login", NewHandler.Login)
	r.POST("/register", NewHandler.Register)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func MaxAllowed(n int) gin.HandlerFunc {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
	}

	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request

		c.Set("sem", sem)
		c.Set("count_request", countReq)

		c.Next()
	}
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
