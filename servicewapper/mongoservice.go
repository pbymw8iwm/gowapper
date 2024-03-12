package servicewapper

import (
	"github.com/astaxie/beego"
	dbwapper "github.com/pbymw8iwm/gowapper/dbwapper"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoManager struct {
	Client *dbwapper.MongoClientWapper //如果不做读写分离，实例化一个client
}

func (p *MongoManager) Stop() error {
	err := p.Client.Close()
	p.Client = nil
	return err
}

func (p *MongoManager) Start(params IServiceParams) error {
	p.Client = new(dbwapper.MongoClientWapper)
	err := p.Client.Connect((params.GetCfg().(*dbwapper.MongoDBParam)))
	if err != nil {
		return err
	}
	return nil
}

func test_mongo() {
	var mongo_mgr *MongoManager
	mongo_mgr = new(MongoManager)
	var mongoparam IServiceParams = &dbwapper.MongoDBParam{
		Dbname:   beego.AppConfig.String("mongodb::db"),
		Uri:      beego.AppConfig.String("mongodb::master_url"),
		RTimeout: 5,
		WTimeout: 5,
	}

	mongo_mgr.Start(mongoparam)
	var a []bson.M
	err := mongo_mgr.Client.FindMany("usercenter", bson.D{{Key: "channel", Value: "android"}}, mongo_mgr.Client.GetFindOpts().SetLimit(10), &a)
	beego.Info("mongo server: master", a, "salve", err) //Bson2Obj
	var b bson.M
	err = mongo_mgr.Client.FindOne("usercenter", bson.D{{Key: "channel", Value: "android"}}, mongo_mgr.Client.GetFindOneOpts().SetProjection(bson.M{"playerid": 1}), &b)
	beego.Info("mongo server: master", b, "salve", err)
	var c []bson.M
	playerid := "503510cd-9fa5-465e-a7cd-9f1e50f6cb07"
	groupStage := []bson.D{
		{{Key: "$match", Value: bson.D{{Key: "playerid", Value: playerid}}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "usercenter"},
			{Key: "let", Value: bson.D{{Key: "friends", Value: "$friends"}}},
			{
				Key: "pipeline",
				Value: bson.A{
					bson.M{
						"$match": bson.M{
							"$expr": bson.M{
								"$or": bson.A{
									bson.M{
										"$in": bson.A{"$playerid", "$$friends.playerid"},
									},
									bson.M{
										"$eq": bson.A{"$playerid", playerid},
									},
								},
							},
						},
					},
				},
			},
			{Key: "as", Value: "userdata"},
		},
		}},
		{{Key: "$project", Value: bson.M{"userdata": 1, "_id": 0}}},
		{{Key: "$unwind", Value: "$userdata"}},
		{{Key: "$replaceRoot", Value: bson.M{"newRoot": bson.M{"$mergeObjects": bson.A{"$userdata", "$$ROOT"}}}}},
		{{Key: "$sort", Value: bson.D{{Key: "decorate", Value: -1}}}},
	}

	err = mongo_mgr.Client.Aggregate("friend", groupStage, &c)
	beego.Info("mongo server: master", b, "salve", err)
}
