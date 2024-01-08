package suggestion

//
// import (
// 	"context"
//
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )
//
// type mongoUserUsage struct {
// 	UserID   string             `bson:"userID"`
// 	Tokens   int                `bson:"tokens"`
// 	Requests int                `bson:"requests"`
// 	UsageID  primitive.ObjectID `bson:"_id"`
// }
//
// type Repository interface{}
//
// type MongoRepository struct {
// 	col *mongo.Collection
// }
//
// func (r MongoRepository) GetSuggestUsage(ctx context.Context, uid string) (Usage, error) {
// 	filter := bson.M{"userID": uid}
//
// 	res := r.col.FindOne(ctx, filter)
// 	if err := res.Err(); err != nil {
// 		return Usage{}, err
// 	}
//
// 	usage := new(Usage)
// 	if err := res.Decode(usage); err != nil {
// 		return Usage{}, err
// 	}
//
// 	return *usage, nil
// }
//
// func (r MongoRepository) DeleteSuggestUsage(ctx context.Context, uid string) error {
// 	filter := bson.M{"userID": uid}
// 	_, err := r.col.DeleteOne(ctx, filter)
// 	return err
// }
//
// func (r MongoRepository) UpsertSuggestUsage(ctx context.Context, uid string, usage SuggestUsage) error {
// 	filter := bson.M{"userID": uid}
// 	update := bson.M{"$add": bson.M{"suggestUsage.tokens": usage.Token, "suggestUsage.requests": 1}}
//
// 	_, err := r.col.UpdateOne(ctx,
// 		filter,
// 		update,
// 		options.Update().SetUpsert(false),
// 	)
// 	if err != nil {
// 		if err != mongo.ErrNoDocuments {
// 			return err
// 		}
// 	}
// 	return r.InsertSuggestUsage(ctx, Usage{UserID: uid, Usage: usage})
// }
//
// func (r MongoRepository) InsertSuggestUsage(ctx context.Context, userUsage Usage) error {
// 	mongoUserUsage := mongoUserUsage{
// 		UserID:   userUsage.UserID,
// 		Tokens:   userUsage.Usage.Token,
// 		Requests: userUsage.Usage.Request,
// 	}
//
// 	_, err := r.col.InsertOne(ctx, mongoUserUsage)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func NewMongoRepository(db *mongo.Collection) (*MongoRepository, error) {
// 	return &MongoRepository{col: db}, nil
// }
