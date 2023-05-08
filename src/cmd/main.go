package main

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/action/command"
	"github.com/cassa10/arq2-tp1/src/domain/action/query"
	"github.com/cassa10/arq2-tp1/src/domain/usecase"
	"github.com/cassa10/arq2-tp1/src/infrastructure/api"
	"github.com/cassa10/arq2-tp1/src/infrastructure/config"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"github.com/cassa10/arq2-tp1/src/infrastructure/repository/mongo"
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

	//repositories
	customerRepo := mongo.NewCustomerRepository(baseLogger, mongoDB, conf.MongoTimeout)
	sellerRepo := mongo.NewSellerRepository(baseLogger, mongoDB, conf.MongoTimeout)
	productRepo := mongo.NewProductRepository(baseLogger, mongoDB, conf.MongoTimeout)
	orderRepo := mongo.NewOrderRepository(baseLogger, mongoDB, conf.MongoTimeout, conf.MongoDatabase)

	//customer
	findCustomerByIdQuery := query.NewFindCustomerById(customerRepo)
	createCustomerCmd := command.NewCreateCustomer(customerRepo)
	updateCustomerCmd := command.NewUpdateCustomer(customerRepo, *findCustomerByIdQuery)
	deleteCustomerCmd := command.NewDeleteCustomer(customerRepo, *findCustomerByIdQuery)

	//seller
	findSellerByIdQuery := query.NewFindSellerById(sellerRepo, productRepo)
	createSellerCmd := command.NewCreateSeller(sellerRepo)
	updateSellerCmd := command.NewUpdateSeller(sellerRepo, *findSellerByIdQuery)
	deleteSellerCmd := command.NewDeleteSeller(sellerRepo, productRepo, *findSellerByIdQuery)

	//product
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

	createOrderUseCase := usecase.NewCreateOrder(baseLogger, *createOrderCmd, *findProductByIdQuery, *findCustomerByIdQuery)
	confirmOrderUseCase := usecase.NewConfirmOrder(baseLogger, *confirmOrderCmd, *findOrderByIdQuery)
	deliveredOrderUseCase := usecase.NewDeliveredOrder(baseLogger, *deliveredOrderCmd, *findOrderByIdQuery)

	app := api.NewApplication(baseLogger, conf, &api.ApplicationUseCases{
		FindCustomerQuery: findCustomerByIdQuery,
		CreateCustomerCmd: createCustomerCmd,
		UpdateCustomerCmd: updateCustomerCmd,
		DeleteCustomerCmd: deleteCustomerCmd,

		FindSellerQuery: findSellerByIdQuery,
		CreateSellerCmd: createSellerCmd,
		UpdateSellerCmd: updateSellerCmd,
		DeleteSellerCmd: deleteSellerCmd,

		FindProductQuery:    findProductByIdQuery,
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
