package main

import (
	"context"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/command"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/action/query"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/usecase"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/api"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/config"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/logger"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/repository/http"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/repository/mongo"
)

func main() {
	conf := config.LoadConfig()
	baseLogger := logger.New(&logger.Config{
		ServiceName:     "arq2-tp1",
		EnvironmentName: conf.Environment,
		LogLevel:        conf.LogLevel,
		LogFormat:       logger.JsonFormat,
	})
	mongoDB := mongo.Connect(context.Background(), baseLogger, conf.MongoURI, conf.MongoDatabase)

	//mongo repositories
	productRepo := mongo.NewProductRepository(baseLogger, mongoDB, conf.MongoTimeout)
	orderRepo := mongo.NewOrderRepository(baseLogger, mongoDB, conf.MongoTimeout, conf.MongoDatabase)

	//http repositories
	sellerRepo := http.NewSellerRepository(baseLogger, http.NewClient(), conf.Seller.UrlFindById)

	//product
	findSellerByIdQuery := query.NewFindSellerById(sellerRepo)
	findProductByIdQuery := query.NewFindProductById(productRepo)
	createProductCmd := command.NewCreateProduct(productRepo, *findSellerByIdQuery)
	updateProductCmd := command.NewUpdateProduct(productRepo, *findProductByIdQuery)
	deleteProductCmd := command.NewDeleteProduct(productRepo, *findProductByIdQuery)
	searchProductQuery := query.NewSearchProducts(productRepo)

	//order
	findOrderByIdQuery := query.NewFindOrderById(orderRepo)
	createOrderCmd := command.NewCreateOrder(orderRepo)
	confirmOrderCmd := command.NewConfirmOrder(orderRepo)
	deliveredOrderCmd := command.NewDeliveredOrder(orderRepo)

	createOrderUseCase := usecase.NewCreateOrder(baseLogger, *createOrderCmd, *findProductByIdQuery)
	confirmOrderUseCase := usecase.NewConfirmOrder(baseLogger, *confirmOrderCmd, *findOrderByIdQuery)
	deliveredOrderUseCase := usecase.NewDeliveredOrder(baseLogger, *deliveredOrderCmd, *findOrderByIdQuery)

	app := api.NewApplication(baseLogger, conf, &api.ApplicationUseCases{
		FindProductQuery:    findProductByIdQuery,
		FindSellerQuery:     findSellerByIdQuery,
		CreateProductCmd:    createProductCmd,
		UpdateProductCmd:    updateProductCmd,
		DeleteProductCmd:    deleteProductCmd,
		SearchProductsQuery: searchProductQuery,

		FindOrderQuery:        findOrderByIdQuery,
		CreateOrderUseCase:    createOrderUseCase,
		ConfirmOrderUseCase:   confirmOrderUseCase,
		DeliveredOrderUseCase: deliveredOrderUseCase,
	})
	baseLogger.Fatal(app.Run())
}
