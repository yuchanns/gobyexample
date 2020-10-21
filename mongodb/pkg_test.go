package mongodb

import (
	"context"
	"github.com/coreos/etcd/pkg/testutil"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestNewMongoClient(t *testing.T) {
	client, err := NewMongoClient()
	testutil.AssertNil(t, err)
	defer client.Close()

	name := "yuchanns"

	cl := client.GetConn().Database("test").Collection("cl1")

	_, err = cl.InsertOne(context.Background(), bson.M{"hello": name})
	testutil.AssertNil(t, err)

	result := struct {
		Hello string
	}{}

	err = cl.FindOne(context.Background(), bson.D{{"hello", name}}).Decode(&result)
	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, result.Hello, name)

	cur, err := cl.Find(context.Background(), bson.D{{"hello", name}})
	testutil.AssertNil(t, err)
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := struct {
			Hello string
		}{}
		err := cur.Decode(&result)
		testutil.AssertNil(t, err)
		testutil.AssertEqual(t, result.Hello, name)
	}
}
