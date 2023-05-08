package mongo

import (
	"context"
	"fmt"
	"github.com/cassa10/arq2-tp1/src/domain/model"
	"github.com/cassa10/arq2-tp1/src/infrastructure/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func handleCloseCursor(cur *mongo.Cursor, ctx context.Context, log model.Logger) {
	err := cur.Close(ctx)
	if err != nil {
		log.WithFields(logger.Fields{"error": err}).Warn("some error when close cursor")
	}
}

func createStringCaseInsensitiveFilter(value string) bson.M {
	return bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf("^%s.*", value), Options: "i"}}
}
