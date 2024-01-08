package suggestion

//
// import (
// 	"context"
// 	"os"
// 	"testing"
// 	"time"
//
// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )
//
// var (
// 	mongoUri = os.Getenv("MONGO_URI")
// 	dbName   = os.Getenv("MONGO_DB") + ".test"
// 	colName  = os.Getenv("MONGO_COL")
// )
//
// func TestRepository(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()
// 	initRepository(t, ctx)
// }
//
// func TestGetUsage(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()
// 	repository := initRepository(t, ctx)
//
// }
//
// func TestUpsertRepository(t *testing.T) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()
// 	repository := initRepository(t, ctx)
//
// 	repository.UpsertSuggestUsage(ctx, "", SuggestUsage{})
// }
//
// func initRepository(t *testing.T, ctx context.Context) *MongoRepository {
// 	opts := options.Client().ApplyURI(mongoUri)
// 	app, err := mongo.Connect(ctx, opts)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, app)
// 	assert.Nil(t, app.Ping(ctx, &readpref.ReadPref{}))
// 	col := app.Database(dbName).Collection(colName)
// 	repository, err := NewMongoRepository(col)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, repository)
// 	return repository
// }
