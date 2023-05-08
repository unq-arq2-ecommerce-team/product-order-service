package mongo

import (
	"context"
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/domain/model/exception"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const customerCollection = "customers"

type customerRepository struct {
	logger  model.Logger
	db      *mongo.Database
	timeout time.Duration
}

func NewCustomerRepository(baseLogger model.Logger, db *mongo.Database, timeout time.Duration) model.CustomerRepository {
	repo := &customerRepository{
		logger:  baseLogger.WithFields(logger.Fields{"logger": "mongo.CustomerRepository", "customerCollection": customerCollection}),
		db:      db,
		timeout: timeout,
	}
	repo.createIndex(context.Background())
	return repo
}

func (r *customerRepository) FindById(ctx context.Context, id int64) (*model.Customer, error) {
	log := r.logger.WithFields(logger.Fields{"method": "FindById", "id": id})
	filter := bson.M{"_id": id}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	var customer model.Customer
	if err := r.db.Collection(customerCollection).FindOne(timeout, filter).Decode(&customer); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.CustomerNotFound{Id: id}
		}
		log.WithFields(logger.Fields{"err": err}).Errorf(fmt.Sprintf("couldn't retrieve documents with filter %s", filter))
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) FindByEmail(ctx context.Context, email string) (*model.Customer, error) {
	log := r.logger.WithFields(logger.Fields{"method": "FindByEmail", "email": email})
	filter := bson.M{"email": email}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	var customer model.Customer
	if err := r.db.Collection(customerCollection).FindOne(timeout, filter).Decode(&customer); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.CustomerNotFound{Email: email}
		}
		log.WithFields(logger.Fields{"err": err}).Errorf(fmt.Sprintf("couldn't retrieve documents with filter %s", filter))
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepository) Create(ctx context.Context, customer model.Customer) (int64, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Create"})
	_, err := r.FindByEmail(ctx, customer.Email)
	if _, customerNotExist := err.(exception.CustomerNotFound); !customerNotExist {
		log.Infof("customer already exist")
		return 0, exception.CustomerAlreadyExist{Email: customer.Email}
	}

	timeoutCtx, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	customerId, err := getNextId(timeoutCtx, log, r.db, r.timeout, customerCollection)
	if err != nil {
		return 0, err
	}
	customer.Id = customerId

	log = log.WithFields(logger.Fields{"customer": customer})
	if _, err := r.db.Collection(customerCollection).InsertOne(timeoutCtx, customer); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Infof("customer already exist")
			return 0, exception.CustomerAlreadyExist{Email: customer.Email}
		}
		log.WithFields(logger.Fields{"err": err}).Error("couldn't create customer")
		return 0, err
	}
	log.Info("customer created successfully")
	return customer.Id, nil
}

func (r *customerRepository) Update(ctx context.Context, customer model.Customer) (bool, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Update", "customerToUpdate": customer})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	updateRes, err := r.db.Collection(customerCollection).UpdateByID(timeout, customer.Id, bson.M{"$set": customer})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Infof("customer with new email already exist")
			return false, exception.CustomerAlreadyExist{Email: customer.Email}
		}
		log.WithFields(logger.Fields{"error": err}).Error("error when update customer")
		return false, err
	}
	if updateRes.ModifiedCount == 0 {
		log.Errorf("couldn't update customer with id %v", customer.Id)
		return false, exception.CustomerCannotUpdate{Id: customer.Id}
	}
	log.Info("customer updated successfully")
	return true, nil
}

func (r *customerRepository) Delete(ctx context.Context, id int64) (bool, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Delete", "customerId": id})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	deleteRes, err := r.db.Collection(customerCollection).DeleteOne(timeout, bson.M{"_id": id})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when delete customer with id %v", id)
		return false, err
	}
	if deleteRes.DeletedCount == 0 {
		log.Infof("customer id %v cannot be deleted", id)
		return false, exception.CustomerCannotDelete{Id: id}
	}
	log.Infof("customer with id %v deleted successfully", id)
	return true, nil
}

func (r *customerRepository) createIndex(ctx context.Context) {
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	index := mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := r.db.Collection(customerCollection).Indexes().CreateOne(timeout, index)
	if err != nil {
		r.logger.WithFields(logger.Fields{"error": err}).Fatalf("could not create mongo index")
	} else {
		r.logger.Infof("mongo index created")
	}
}
