package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/domain/model/exception"
	"github.com/unq-arq2-ecommerce-team/products-orders-service/src/infrastructure/logger"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type sellerRepository struct {
	logger      model.Logger
	client      *http.Client
	findByIdUrl string
}

func NewSellerRepository(baselogger model.Logger, client *http.Client, findByIdUrl string) model.SellerRepository {
	return sellerRepository{
		logger:      baselogger.WithFields(logger.Fields{"logger": "http.SellerRepository"}),
		client:      client,
		findByIdUrl: findByIdUrl,
	}
}

func (repo sellerRepository) FindById(ctx context.Context, sellerId int64) (*model.Seller, error) {
	log := repo.logger.WithFields(logger.Fields{"findByIdUrl": repo.findByIdUrl, "sellerId": sellerId})
	log.Debugf("http find seller by id")
	url := strings.Replace(repo.findByIdUrl, "{sellerId}", strconv.FormatInt(sellerId, 10), -1)
	log = repo.logger.WithFields(logger.Fields{"url": url})
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("error when create request")
		return nil, err
	}
	sw := time.Now()
	res, err := repo.client.Do(req)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Error("http error post add workflow event")
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	log.Debugf("request find by id seller finished in %s", time.Since(sw))

	switch statusCode := res.StatusCode; {
	case statusCode >= 200 && statusCode <= 299:
		rawBody, _ := io.ReadAll(res.Body)
		log = log.WithFields(logger.Fields{"bodyRaw": rawBody})
		var seller model.Seller
		err = json.Unmarshal(rawBody, &seller)
		return &seller, nil
	case statusCode == http.StatusNotFound:
		return nil, exception.SellerNotFound{Id: sellerId}
	default:
		return nil, fmt.Errorf("unexpected error with status code %v when find seller by id from seller repository with url %s", statusCode, url)
	}
}
