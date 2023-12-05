package xmongo

import (
	"context"
	"errors"
	"fmt"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	cfg *pb.Jmongo
	db  *mongo.Database
}

var cli *MongoClient

func NewMongoClient(cfg *pb.Jmongo) *MongoClient {
	// if cli != nil {
	// 	return cli
	// }
	cli := &MongoClient{
		cfg: cfg,
	}
	cli.connectDB()
	return cli
}

func (t *MongoClient) connectDB() {
	uri := fmt.Sprintf("mongodb://%s/?maxPoolSize=%d", t.cfg.Address, t.cfg.MaxPoolSize)
	clientOptions := options.Client().ApplyURI(uri).SetConnectTimeout(time.Duration(t.cfg.Timeout) * time.Second)
	if t.cfg.Username != "" && t.cfg.Password != "" {
		clientOptions = clientOptions.SetAuth(options.Credential{
			AuthMechanism: "SCRAM-SHA-256",
			Username:      t.cfg.Username,
			Password:      t.cfg.Password,
		})
	}
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	t.db = client.Database(t.cfg.Db)
}

func (t *MongoClient) Disconnect() {
	t.db.Client().Disconnect(context.TODO())
}

func (t *MongoClient) Exec(fn func(*mongo.Database) error) error {
	return fn(t.db)
}

// 查询所有已经创建的index
func getIndexNames(coll *mongo.Collection) (map[string]bool, error) {
	cursor, err := coll.Indexes().List(context.TODO())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	names := map[string]bool{}
	for i := 0; i < len(result); i++ {
		for k, v := range result[i] {
			if k == "name" {
				names[v.(string)] = true
			}
		}
	}
	return names, nil
}

func checkInit() error {
	if cli == nil {
		return errors.New(utils.Const_Mongo_Not_Instance)
	}
	if cli.db == nil {
		return errors.New(utils.Const_Mongo_Not_Connect)
	}
	return nil
}

// 创建Index
func CreateIndex(collName string, indexs map[string]mongo.IndexModel) error {
	if e := checkInit(); e != nil {
		return e
	}

	//查询已经创建的索引
	coll := cli.db.Collection(collName)
	names, err := getIndexNames(coll)
	if err != nil {
		return err
	}
	//检测索引是否创建，未创建就创建新的
	for k, v := range indexs {
		if _, ok := names[k]; ok {
			continue //已经创建
		}
		_, err := coll.Indexes().CreateOne(context.TODO(), v)
		if err != nil {
			return err
		}
	}
	return nil
}

// 同步执行
func ExecSync(fn func(*mongo.Database) error) error {
	if e := checkInit(); e != nil {
		return e
	}
	return fn(cli.db)
}
