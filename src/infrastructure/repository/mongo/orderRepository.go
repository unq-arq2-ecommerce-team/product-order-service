package mongo

import (
	"context"
	"fmt"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/dto"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const orderCollection = "orders"

type orderRepository struct {
	logger   model.Logger
	db       *mongo.Database
	database string
	timeout  time.Duration
}

func NewOrderRepository(baseLogger model.Logger, db *mongo.Database, timeout time.Duration, database string) model.OrderRepository {
	repo := &orderRepository{
		logger:   baseLogger.WithFields(logger.Fields{"logger": "mongo.OrderRepository", "orderCollection": orderCollection}),
		db:       db,
		timeout:  timeout,
		database: database,
	}
	repo.createIndexes(context.Background())
	return repo
}

func (r *orderRepository) Create(ctx context.Context, order model.Order) (int64, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "Create", "order": order})
	timeoutCtx, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	client := r.db.Client()

	//transaction
	session, err := client.StartSession()
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when start session")
		return 0, err
	}
	defer session.EndSession(timeoutCtx)

	err = session.StartTransaction()
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when start transaction")
		return 0, err
	}
	if err = mongo.WithSession(timeoutCtx, session, func(sc mongo.SessionContext) error {
		db := sc.Client().Database(r.database)
		productFilter := bson.M{"_id": order.GetProductId(), "stock": bson.M{"$gte": 1}}
		updateProductRes, err := db.Collection(productCollection).UpdateOne(sc, productFilter, bson.M{"$inc": bson.M{"stock": -1}})
		if err != nil || updateProductRes.ModifiedCount == 0 {
			_ = session.AbortTransaction(sc)
			log.WithFields(logger.Fields{"error": err}).Error("some error found or product with stock not found")
			return exception.ProductWithNoStock{Id: order.GetProductId()}
		}
		orderId, err := getNextId(sc, log, db, r.timeout, orderCollection)
		if err != nil {
			_ = session.AbortTransaction(sc)
			return err
		}
		order.Id = orderId
		orderDTO := dto.NewOrderDTOFrom(order)
		log = log.WithFields(logger.Fields{"orderDTO": orderDTO})
		if _, err := db.Collection(orderCollection).InsertOne(sc, orderDTO); err != nil {
			_ = session.AbortTransaction(sc)
			log.WithFields(logger.Fields{"error": err}).Error("couldn't create order")
			return err
		}

		err = session.CommitTransaction(sc)
		if err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("error when commit transaction")
			return err
		}
		return nil
	}); err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when create order mongo session")
		return 0, err
	}

	log.Info("order created successfully")
	return order.Id, nil
}

func (r *orderRepository) FindById(ctx context.Context, id int64) (*model.Order, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "FindById", "id": id})
	filter := bson.M{"_id": id}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	var orderDTO dto.OrderDTO
	if err := r.db.Collection(orderCollection).FindOne(timeout, filter).Decode(&orderDTO); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.OrderNotFound{Id: id}
		}
		log.WithFields(logger.Fields{"error": err}).Errorf(fmt.Sprintf("couldn't retrieve documents with filter %s", filter))
		return nil, err
	}
	order, err := orderDTO.Map()
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when map order dto to model order")
		return nil, err
	}
	log.Debug("order found")
	return &order, nil
}

func (r *orderRepository) Update(ctx context.Context, order model.Order) (bool, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "Update", "orderToUpdate": order})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()

	orderDTO := dto.NewOrderDTOFrom(order)
	updateRes, err := r.db.Collection(orderCollection).UpdateByID(timeout, orderDTO.Id, bson.M{"$set": orderDTO})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when update order")
		return false, err
	}
	if updateRes.ModifiedCount == 0 {
		log.Errorf("couldn't update order with id %v", order.Id)
		return false, exception.OrderCannotUpdate{Id: order.Id}
	}
	log.Info("order updated successfully")
	return true, nil
}

func (r *orderRepository) createIndexes(ctx context.Context) {
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	indexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{"createdOn", -1}},
			Options: options.Index().SetUnique(false),
		},
		{
			Keys:    bson.D{{"deliveryDate", -1}},
			Options: options.Index().SetUnique(false),
		},
	}
	_, err := r.db.Collection(orderCollection).Indexes().CreateMany(timeout, indexes)
	if err != nil {
		r.logger.WithFields(logger.Fields{"error": err}).Fatalf("could not create mongo index")
	} else {
		r.logger.Infof("mongo index created")
	}
}
