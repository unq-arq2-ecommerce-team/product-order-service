package api

import (
	"fmt"
	swaggerDocs "github.com/cassa10/arq2-tp1/docs"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/usecase"
	v1 "github.com/cassa10/arq2-tp1/src/infrastructure/api/v1"
	"github.com/cassa10/arq2-tp1/src/infrastructure/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"net/http"
)

// Application
// @title arq2-tp1 API
// @version 1.0
// @description api for tp arq2-tp1
// @contact.name API SUPPORT
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /
// @query.collection.format multi
type Application interface {
	Run() error
}

type application struct {
	logger model.Logger
	config config.Config
	*ApplicationUseCases
}

type ApplicationUseCases struct {
	//customer
	CreateCustomerCmd *command.CreateCustomer
	UpdateCustomerCmd *command.UpdateCustomer
	DeleteCustomerCmd *command.DeleteCustomer
	FindCustomerQuery *query.FindCustomerById
	//seller
	CreateSellerCmd *command.CreateSeller
	UpdateSellerCmd *command.UpdateSeller
	DeleteSellerCmd *command.DeleteSeller
	FindSellerQuery *query.FindSellerById
	//product
	CreateProductCmd    *command.CreateProduct
	UpdateProductCmd    *command.UpdateProduct
	DeleteProductCmd    *command.DeleteProduct
	FindProductQuery    *query.FindProductById
	SearchProductsQuery *query.SearchProducts
	//order
	FindOrderQuery        *query.FindOrderById
	CreateOrderUseCase    *usecase.CreateOrder
	ConfirmOrderUseCase   *usecase.ConfirmOrder
	DeliveredOrderUseCase *usecase.DeliveredOrder
}

func NewApplication(l model.Logger, conf config.Config, applicationUseCases *ApplicationUseCases) Application {
	return &application{
		logger:              l,
		config:              conf,
		ApplicationUseCases: applicationUseCases,
	}
}

func (app *application) Run() error {
	swaggerDocs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", app.config.Port)

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	router := gin.Default()
	router.GET("/", HealthCheck)

	rv1 := router.Group("/api/v1")
	{
		rv1Customer := rv1.Group("/customer")
		rv1Customer.POST("", v1.CreateCustomerHandler(app.logger, app.CreateCustomerCmd))
		rv1Customer.GET("/:customerId", v1.FindCustomerHandler(app.logger, app.FindCustomerQuery))
		rv1Customer.DELETE("/:customerId", v1.DeleteCustomerHandler(app.logger, app.DeleteCustomerCmd))
		rv1Customer.PUT("/:customerId", v1.UpdateCustomerHandler(app.logger, app.UpdateCustomerCmd))
	}
	{
		rv1Seller := rv1.Group("/seller")
		rv1Seller.POST("", v1.CreateSellerHandler(app.logger, app.CreateSellerCmd))
		rv1Seller.GET("/:sellerId", v1.FindSellerHandler(app.logger, app.FindSellerQuery))
		rv1Seller.DELETE("/:sellerId", v1.DeleteSellerHandler(app.logger, app.DeleteSellerCmd))
		rv1Seller.PUT("/:sellerId", v1.UpdateSellerHandler(app.logger, app.UpdateSellerCmd))
	}
	{
		rv1.POST("/seller/:sellerId/product", v1.CreateProductHandler(app.logger, app.CreateProductCmd))
		rv1Product := rv1.Group("/seller/product")
		rv1Product.GET("/:productId", v1.FindProductHandler(app.logger, app.FindProductQuery))
		rv1Product.DELETE("/:productId", v1.DeleteProductHandler(app.logger, app.DeleteProductCmd))
		rv1Product.PUT("/:productId", v1.UpdateProductHandler(app.logger, app.UpdateProductCmd))
		rv1Product.GET("/search", v1.SearchProductsHandler(app.logger, app.SearchProductsQuery))
	}
	{
		rv1Order := rv1.Group("/order")
		rv1Order.POST("", v1.CreateOrderHandler(app.logger, app.CreateOrderUseCase))
		rv1Order.GET("/:orderId", v1.FindOrderHandler(app.logger, app.FindOrderQuery))
		rv1Order.POST("/:orderId/confirm", v1.ConfirmOrderHandler(app.logger, app.ConfirmOrderUseCase))
		rv1Order.POST("/:orderId/delivered", v1.DeliveredOrderHandler(app.logger, app.DeliveredOrderUseCase))
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.logger.Infof("running http server on port %d", app.config.Port)
	return router.Run(fmt.Sprintf(":%d", app.config.Port))
}

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health check
// @Accept */*
// @Produce json
// @Success 200 {object} HealthCheckRes
// @Router / [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthCheckRes{Data: "Server is up and running"})
}

type HealthCheckRes struct {
	Data string `json:"data"`
}
