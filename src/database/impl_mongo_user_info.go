package database

import (
	"context"
	"rosenchat/src/configs"
	"rosenchat/src/database/mongodb"
	"rosenchat/src/exception"
	"rosenchat/src/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var conf = configs.Get()
var log = logger.Get()

// implMongoUserInfo implements the IUserInfo interface using MongoDB.
type implMongoUserInfo struct {
	client *mongo.Client
}

func (i *implMongoUserInfo) GetUserInfo(userID string) (*UserInfoDTO, error) {
	timeoutDuration := time.Duration(conf.Mongo.OperationTimeoutSec) * time.Second
	dbCallCtx, cancelFunc := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancelFunc()

	result := i.getColl().FindOne(dbCallCtx, bson.M{"_id": userID})
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, exception.UserNotFound()
		}

		log.Errorf("Unexpected error in FindOne call: %+v", result.Err())
		return nil, result.Err()
	}

	userInfo := &UserInfoDTO{}
	if err := result.Decode(userInfo); err != nil {
		log.Errorf("Unexpected error while decoding the user document: %+v", err)
		return nil, err
	}

	return userInfo, nil
}

func (i *implMongoUserInfo) PutUserInfo(info *UserInfoDTO) error {
	filter := bson.M{"_id": info.ID}
	update := bson.M{"$set": info}
	opts := options.Update().SetUpsert(true)

	timeoutDuration := time.Duration(conf.Mongo.OperationTimeoutSec) * time.Second
	dbCallCtx, cancelFunc := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancelFunc()

	if _, err := i.getColl().UpdateOne(dbCallCtx, filter, update, opts); err != nil {
		log.Errorf("Unexpected error in FindOneAndUpdate call: %+v", err)
		return err
	}

	return nil
}

func (i *implMongoUserInfo) init() {
	i.client = mongodb.GetClient()

	if err := i.addIndexes(); err != nil {
		panic("failed to add indexes for " + userInfoTableName + " because: " + err.Error())
	}
	log.Infof("Indexes added for collection: %s", userInfoTableName)
}

// addIndexes adds the required indexes to the database if not already present.
func (i *implMongoUserInfo) addIndexes() error {
	return nil
}

func (i *implMongoUserInfo) getColl() *mongo.Collection {
	return i.client.Database(conf.Mongo.DatabaseName).Collection(userInfoTableName)
}
