package mongo

import (
	"context"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/dto"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const counterCollection = "counters"

// getNextId log errors
func getNextId(parentCtx context.Context, baseLogger model.Logger, db *mongo.Database, timeoutDuration time.Duration, collection string) (int64, error) {
	log := baseLogger.WithFields(logger.Fields{"method": "getNextId", "collection of _id": collection})
	opts := options.RunCmd().SetReadPreference(readpref.Primary())
	command := bson.D{
		{"findAndModify", counterCollection},
		{"query", bson.D{{"_id", collection}}},
		{"update", bson.D{{"$inc", bson.D{{"seq", 1}}}}},
		{"new", true},
	}
	var res dto.NextIdResponse
	if err := db.RunCommand(parentCtx, command, opts).Decode(&res); err != nil {
		log.WithFields(logger.Fields{"error": err}).Errorf("get next id error with counter collection %s and _id %s", counterCollection, collection)
		return 0, err
	}
	log.Debugf("get next id successful with value %v", res.Value.Seq)
	return res.Value.Seq, nil
}
