package mongo

import (
	"context"
	"fmt"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"time"
)

const productCollection = "products"

type productRepository struct {
	logger  model.Logger
	db      *mongo.Database
	timeout time.Duration
}

func NewProductRepository(baseLogger model.Logger, db *mongo.Database, timeout time.Duration) model.ProductRepository {
	repo := &productRepository{
		logger:  baseLogger.WithFields(logger.Fields{"logger": "mongo.ProductRepository", "productCollection": productCollection}),
		db:      db,
		timeout: timeout,
	}
	repo.createIndexes(context.Background())
	return repo
}

func (r *productRepository) FindById(ctx context.Context, id int64) (*model.Product, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "FindById", "id": id})
	filter := bson.M{"_id": id}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	var product model.Product
	if err := r.db.Collection(productCollection).FindOne(timeout, filter).Decode(&product); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.ProductNotFound{Id: id}
		}
		log.WithFields(logger.Fields{"err": err}).Errorf(fmt.Sprintf("couldn't retrieve documents with filter %s", filter))
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, product model.Product) (bool, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "Update", "productToUpdate": product})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	updateRes, err := r.db.Collection(productCollection).UpdateByID(timeout, product.Id, bson.M{"$set": product})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when update product")
		return false, err
	}
	if updateRes.ModifiedCount == 0 {
		log.Errorf("couldn't update product with id %v", product.Id)
		return false, exception.ProductCannotUpdate{Id: product.Id}
	}
	log.Info("product updated successfully")
	return true, nil
}

func (r *productRepository) Delete(ctx context.Context, id int64) (bool, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "Delete", "productId": id})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	deleteRes, err := r.db.Collection(productCollection).DeleteOne(timeout, bson.M{"_id": id})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when delete product with id %v", id)
		return false, err
	}
	if deleteRes.DeletedCount == 0 {
		log.Infof("product id %v cannot be deleted", id)
		return false, exception.ProductCannotDelete{Id: id}
	}
	log.Infof("product with id %v deleted successfully", id)
	return true, nil
}

func (r *productRepository) DeleteAllBySellerId(ctx context.Context, sellerId int64) (bool, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "DeleteAllBySellerId", "sellerId": sellerId})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	deleteRes, err := r.db.Collection(productCollection).DeleteMany(timeout, bson.M{"sellerId": sellerId})
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when delete product with sellerId %v", sellerId)
		return false, err
	}
	log.Infof("%v products was deleted with sellerId %v", deleteRes.DeletedCount, sellerId)
	return true, nil
}

func (r *productRepository) Create(ctx context.Context, product model.Product) (int64, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "Create"})
	timeoutCtx, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()

	productId, err := getNextId(timeoutCtx, log, r.db, r.timeout, productCollection)
	if err != nil {
		return 0, err
	}
	product.Id = productId

	log = log.WithFields(logger.Fields{"product": product})
	if _, err := r.db.Collection(productCollection).InsertOne(timeoutCtx, product); err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("couldn't create product")
		return 0, err
	}
	log.Info("product created successfully")
	return product.Id, nil
}

func (r *productRepository) FindAllBySellerId(ctx context.Context, sellerId int64) ([]model.Product, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"method": "FindById", "sellerId": sellerId})
	filter := bson.M{"sellerId": sellerId}
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()

	cur, err := r.db.Collection(productCollection).Find(timeout, filter)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf(fmt.Sprintf("something went wrong when find with filter %s", filter))
		return nil, err
	}
	defer handleCloseCursor(cur, ctx, log)
	products := make([]model.Product, 0)
	for cur.Next(timeout) {
		var product model.Product
		if err := cur.Decode(&product); err != nil {
			log.WithFields(logger.Fields{"error": err}).Errorf(fmt.Sprintf("something went wrong when want decode product"))
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) Search(ctx context.Context, searchFilters model.ProductSearchFilter, pagingReq model.PagingRequest) ([]model.Product, model.Paging, error) {
	log := r.logger.WithRequestId(ctx).WithFields(logger.Fields{"searchFilters": searchFilters, "paging": pagingReq})
	filter := getFilter(searchFilters)
	log = log.WithFields(logger.Fields{"mongoFields": filter})
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()

	products := make([]model.Product, 0)
	total, err := r.db.Collection(productCollection).CountDocuments(ctx, filter)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("error when count documents for paging")
		return products, model.NewEmptyPage(), err
	}
	totalPages := int(math.Ceil(float64(total) / float64(pagingReq.Size)))
	skip := pagingReq.Size * (pagingReq.Page)
	limit := pagingReq.Size

	emptyPage := model.NewEmptyPage()
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))
	cur, err := r.db.Collection(productCollection).Find(timeout, filter, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Infof("successful search - no documents found")
			return products, emptyPage, nil
		}
		log.WithFields(logger.Fields{"error": err}).Error("error when find documents")
		return products, emptyPage, err
	}
	defer handleCloseCursor(cur, ctx, log)

	for cur.Next(timeout) {
		var product model.Product
		if err := cur.Decode(&product); err != nil {
			log.WithFields(logger.Fields{"error": err}).Error("error when decode some product")
			return []model.Product{}, emptyPage, err
		}
		products = append(products, product)
	}

	pagingResult := model.NewPaging(int(total), len(products), totalPages, pagingReq.Page)
	log.Infof("successful found products with pagingResult %s", pagingResult)
	return products, pagingResult, nil
}

func (r *productRepository) createIndexes(ctx context.Context) {
	timeout, cf := context.WithTimeout(ctx, r.timeout)
	defer cf()
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{"sellerId", 1}},
		},
		{
			Keys: bson.D{
				{"price", 1},
				{"category", 1},
				{"stock", -1},
			},
		},
	}
	_, err := r.db.Collection(productCollection).Indexes().CreateMany(timeout, indexes)
	if err != nil {
		r.logger.WithFields(logger.Fields{"error": err}).Fatalf("could not create mongo index")
	} else {
		r.logger.Infof("mongo index created")
	}
}

// bson.M{"$regex": bson.RegEx{Pattern: id, Options: "i"}}
func getFilter(query model.ProductSearchFilter) bson.M {
	return NewFilterBuilder().
		AppendAndOpFilterIf(query.Name != "", bson.M{"name": createStringCaseInsensitiveFilter(query.Name)}).
		AppendAndOpFilterIf(query.Category != "", bson.M{"category": createStringCaseInsensitiveFilter(query.Category)}).
		AppendAndOpFilterIf(query.SellerId > 0, bson.M{"sellerId": query.SellerId}).
		AppendAndOpFilterIf(query.ContainsAnyPriceFilter(), bson.M{"price": bson.M{mongoGTEOp: query.GetPriceMinOrDefault(), mongoLTEOp: query.GetPriceMaxOrDefault()}}).
		Build()
}
