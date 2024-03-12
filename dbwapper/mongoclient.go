package dbwapper

import (
	"context"
	"fmt"

	//"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDBParam struct {
	Dbname   string
	Uri      string //mongodb://root:****@s-j6c8f8e0198e0fc4.mongodb.rds.aliyuncs.com:3717,s-j6ca42972cb2af44.mongodb.rds.aliyuncs.com:3717/admin
	RTimeout int32  //读操作超时时间
	WTimeout int32  //写操作超时时间
}

//参数分别是 数据库名 和 uri （mongodb://用户名:密码@主机ip:端口）
func (p *MongoDBParam) GetCfg() interface{} {
	return p
}

type MongoClientWapper struct {
	Database *mongo.Database
	Client   *mongo.Client
	RTimeout time.Duration //读超时时间
	WTimeout time.Duration //写超时时间 *time.Second
}

func (this *MongoClientWapper) GetCollection(col string) *mongo.Collection {
	copy_col, err := this.Database.Collection(col).Clone()
	if err != nil {
		beego.Critical("copy collection:", col, "error:", err)
		panic(err)
	}
	return copy_col
}

func (this *MongoClientWapper) CreateIndexs(collection string, keys bson.D, unique bool) {
	opts := options.CreateIndexes().SetMaxTime(this.WTimeout)
	col := this.GetCollection(collection)
	indexView := col.Indexes()

	// 创建索引
	result, err := indexView.CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    keys,
			Options: options.Index().SetUnique(unique).SetBackground(true).SetSparse(true),
		},
		opts,
	)
	if err != nil {
		beego.Informational("创建索引结果：", err, result)
	}
}

func (this *MongoClientWapper) DropIndex(collection string, key string) {
	opts := options.DropIndexes().SetMaxTime(this.WTimeout * time.Second)
	col := this.GetCollection(collection)
	indexView := col.Indexes()
	// 创建索引
	result, err := indexView.DropOne(
		context.Background(),
		key, opts,
	)
	beego.Informational("删除索引结果：", err, result)
}

func (this *MongoClientWapper) Close() error {
	return this.Client.Disconnect(context.Background())
}

func (this *MongoClientWapper) Connect(param *MongoDBParam) (err error) {
	this.RTimeout = time.Duration(param.RTimeout) * time.Second
	this.WTimeout = time.Duration(param.WTimeout) * time.Second

	ctx, _ := context.WithTimeout(context.Background(), this.RTimeout)
	clientOptions := options.Client().ApplyURI(param.Uri)
	clientOptions.SetDirect(true)
	clientOptions.SetMaxPoolSize(1024)
	clientOptions.SetConnectTimeout(this.RTimeout)
	// 连接到MongoDB
	this.Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		beego.Error("数据库连接失败", param.Uri)
		return err
	}
	this.Database = this.Client.Database(param.Dbname)
	if err = this.Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		beego.Error("数据库连接失败", param.Uri)
		return err
	}
	//this.RWLock = &sync.RWMutex{}
	beego.Informational("数据库连接成功", param.Uri)
	return nil
}
func (this *MongoClientWapper) UpdateOne(collection string, filter interface{}, updater interface{}, opt *options.UpdateOptions) error {
	ctx, _ := context.WithTimeout(context.Background(), this.WTimeout)
	result, err := this.GetCollection(collection).UpdateOne(ctx, filter, updater, opt)
	if result.MatchedCount != 0 {
		beego.Informational("matched and replaced an existing document")
		return err
	}
	if result.UpsertedCount != 0 {
		beego.Informational("inserted a new document with ID %v\n", result.UpsertedID)
	}
	return err
}
func (this *MongoClientWapper) UpdateMany(collection string, filter interface{}, updater interface{}, opt *options.UpdateOptions) error {
	ctx, _ := context.WithTimeout(context.Background(), this.WTimeout)
	result, err := this.GetCollection(collection).UpdateMany(ctx, filter, updater, opt)
	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")
		return err
	}
	return err
}

func (this *MongoClientWapper) FindOneAndUpdate(collection string, filter interface{}, updater interface{}, opt *options.FindOneAndUpdateOptions, out interface{}) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), this.WTimeout)
	//opts := options.FindOneAndUpdate().SetUpsert(true)
	err = this.GetCollection(collection).FindOneAndUpdate(ctx, filter, updater, opt).Decode(out)
	logs.Info("UpdateOne:", err)
	return
}

func (this *MongoClientWapper) GetAutoIncrId(collection string, Type string, Incr int64) (ID int64, err error) {
	type AutoIncrUserId struct {
		Id int64 `json:"id,omitempty"`
	}
	var IdInfo AutoIncrUserId
	err = this.FindOneAndUpdate(collection,
		bson.D{{Key: "type", Value: Type}},
		bson.M{"$inc": bson.M{"id": int64(Incr)}},
		options.FindOneAndUpdate().SetUpsert(true),
		&IdInfo)
	if err != nil {
		ID = IdInfo.Id
	}
	return
}

func (this *MongoClientWapper) Count(collection string, filter interface{}, opt *options.CountOptions) (count int64, err error) {
	ctx, _ := context.WithTimeout(context.Background(), this.RTimeout)
	count, err = this.GetCollection(collection).CountDocuments(
		ctx,
		filter,
		opt)

	if err != nil {
		beego.Error("Count failed err:%v", err)
	}
	return
}
func (this *MongoClientWapper) Bson2Obj(val interface{}, obj interface{}) error {
	data, err := bson.Marshal(val)
	if err != nil {
		return err
	}
	bson.Unmarshal(data, obj)
	return nil
}
func (this *MongoClientWapper) GetFindOpts() *options.FindOptions {
	return options.Find()
}

func (this *MongoClientWapper) GetFindOneOpts() *options.FindOneOptions {
	return options.FindOne()
}

func (this *MongoClientWapper) FindOne(collection string, filter interface{}, opt *options.FindOneOptions, out interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), this.RTimeout)

	err := this.GetCollection(collection).FindOne(ctx, filter, opt).Decode(out)
	if err != nil {
		return err
	}
	return nil
}

func (this *MongoClientWapper) FindMany(collection string, filter interface{}, opt *options.FindOptions, out interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), this.RTimeout)
	cursor, err := this.GetCollection(collection).Find(ctx, filter, opt)
	if err != nil {
		beego.Error("Count failed err:%v", err)
		return err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(context.TODO(), out); err != nil {
		beego.Error(err)
	}
	return nil
}

func (this *MongoClientWapper) Aggregate(collection string, pipelines interface{}, out *[]bson.M) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), this.RTimeout)
	opts := options.Aggregate().SetMaxTime(this.RTimeout)
	cursor, err := this.GetCollection(collection).Aggregate(
		ctx,
		pipelines,
		opts)
	if err != nil {
		beego.Error(err) // A pipeline stage specification object must contain exactly one field
	}
	if err = cursor.All(context.TODO(), out); err != nil {
		beego.Error(err)
	}
	return err
}
