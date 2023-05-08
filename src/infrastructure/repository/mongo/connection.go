package mongo

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect(ctx context.Context, baseLogger model.Logger, uri, database string) *mongo.Database {

	log := baseLogger.WithFields(logger.Fields{"logger": "mongo", "database": database})
	ctx, cf := context.WithTimeout(ctx, 10*time.Second)
	defer cf()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Fatalf("an error has occurred while trying to connect to mongo cluster")
	}

	// check the connection
	if err = client.Ping(ctx, nil); err != nil {
		log.WithFields(logger.Fields{"error": err}).Fatalf("could not connect to mongo cluster")
	}

	log.Info("successfully connected to mongo cluster")
	return client.Database(database)
}
