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

const sellerCollection = "sellers"

type sellerRepository struct {
	logger  model.Logger
	db      *mongo.Database
	timeout time.Duration
}

func NewSellerRepository(baseLogger model.Logger, db *mongo.Database, timeout time.Duration) model.SellerRepository {
	repo := &sellerRepository{
		logger:  baseLogger.WithFields(logger.Fields{"logger": "mongo.SellerRepository", "sellerCollection": sellerCollection}),
		db:      db,
		timeout: timeout,
	}
	repo.createIndex(context.Background())
	return repo
}

func (r *sellerRepository) FindById(ctx context.Context, id int64) (*model.Seller, error) {
	log := r.logger.WithFields(logger.Fields{"method": "FindById", "id": id})
	filter := bson.M{"_id": id}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	var seller model.Seller
	if err := r.db.Collection(sellerCollection).FindOne(timeout, filter).Decode(&seller); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.SellerNotFound{Id: id}
		}
		log.WithFields(logger.Fields{"error": err}).Errorf(fmt.Sprintf("couldn't retrieve documents with filter %s", filter))
		return nil, err
	}
	return &seller, nil
}

func (r *sellerRepository) FindByName(ctx context.Context, name string) (*model.Seller, error) {
	log := r.logger.WithFields(logger.Fields{"method": "FindByEmail", "name": name})
	filter := bson.M{"name": name}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	var seller model.Seller
	if err := r.db.Collection(sellerCollection).FindOne(timeout, filter).Decode(&seller); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.SellerNotFound{Name: name}
		}
		log.WithFields(logger.Fields{"error": err}).Errorf(fmt.Sprintf("couldn't retrieve documents with filter %s", filter))
		return nil, err
	}
	return &seller, nil
}

func (r *sellerRepository) Create(ctx context.Context, seller model.Seller) (int64, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Create"})

	_, err := r.FindByName(ctx, seller.Name)
	if _, sellerNotExist := err.(exception.SellerNotFound); !sellerNotExist {
		log.Infof("seller already exist")
		return 0, exception.SellerAlreadyExist{Name: seller.Name}
	}

	timeoutCtx, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	sellerId, err := getNextId(timeoutCtx, log, r.db, r.timeout, sellerCollection)
	if err != nil {
		return 0, err
	}
	seller.Id = sellerId
	log = log.WithFields(logger.Fields{"seller": seller})
	if _, err := r.db.Collection(sellerCollection).InsertOne(timeoutCtx, seller); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Infof("seller already exist")
			return 0, exception.SellerAlreadyExist{Name: seller.Name}
		}
		log.WithFields(logger.Fields{"error": err}).Error("couldn't create seller")
		return 0, err
	}
	log.Info("seller created successfully")
	return seller.Id, nil
}

func (r *sellerRepository) Update(ctx context.Context, seller model.Seller) (bool, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Update", "sellerToUpdate": seller})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	updateRes, err := r.db.Collection(sellerCollection).UpdateByID(timeout, seller.Id, bson.M{"$set": seller})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Infof("seller with new name already exist")
			return false, exception.SellerAlreadyExist{Name: seller.Name}
		}
		log.WithFields(logger.Fields{"error": err}).Error("couldn't update seller")
		return false, err
	}
	if updateRes.ModifiedCount == 0 {
		log.Errorf("couldn't update seller with id %v", seller.Id)
		return false, exception.SellerCannotUpdate{Id: seller.Id}
	}
	log.Info("seller updated successfully")
	return true, nil
}

func (r *sellerRepository) Delete(ctx context.Context, id int64) (bool, error) {
	log := r.logger.WithFields(logger.Fields{"method": "Delete", "sellerId": id})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	deleteRes, err := r.db.Collection(sellerCollection).DeleteOne(timeout, bson.M{"_id": id})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when delete seller with id %v", id)
		return false, err
	}
	if deleteRes.DeletedCount == 0 {
		log.Infof("seller id %v cannot be deleted", id)
		return false, exception.SellerCannotDelete{Id: id}
	}
	log.Infof("seller with id %v deleted successfully", id)
	return true, nil
}

func (r *sellerRepository) createIndex(ctx context.Context) {
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	index := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := r.db.Collection(sellerCollection).Indexes().CreateOne(timeout, index)
	if err != nil {
		r.logger.WithFields(logger.Fields{"error": err}).Fatalf("could not create mongo index")
	} else {
		r.logger.Infof("mongo index created")
	}
}
