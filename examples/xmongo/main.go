package main

import (
	"context"
	"lark/com/pkgs/xloadcfg"
	"lark/com/pkgs/xlog"
	"lark/com/pkgs/xmongo"
	"lark/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Others    string `json:"others"`
}

// type ActInfoDB struct {
// 	ActID uint32
// 	Act   *pb.MsgActInfo
// 	Data  *pb.MsgActGlobleData
// }

// func loadRLs(cfg *pb.Jmongo) map[uint32]map[uint32]uint32 {
// 	cfg.Db = "yw_czs_cn1_center1"
// 	cli := xmongo.NewMongoClient(cfg)
// 	dat := []*pb.MsgRankListData{}
// 	cli.Exec(func(db *mongo.Database) error {
// 		cur, err := db.Collection("ranklist").Find(context.TODO(), bson.M{"type": 10})
// 		cur.All(context.TODO(), &dat)
// 		return err
// 	})
// 	cli.Disconnect()
// 	// dp := map[uint32]map[uint32]uint32{}
// 	// for _, rl := range dat {
// 	// 	for i, cell := range rl.Rank {
// 	// 		sid := cell.Id >> 16
// 	// 		mp, ok := dp[uint32(sid)]
// 	// 		if !ok || mp == nil {
// 	// 			mp = map[uint32]uint32{}
// 	// 			dp[uint32(sid)] = mp
// 	// 		}
// 	// 		mp[uint32(i+1)] = uint32(cell.Id)
// 	// 	}
// 	// }
// 	// return dp
// 	save := []*pb.MsgActUnionTaskDrawRecover{}
// 	for _, rl := range dat {
// 		for i, cell := range rl.Rank {
// 			save = append(save, &pb.MsgActUnionTaskDrawRecover{
// 				Unionid: uint32(cell.Id),
// 				Rank:    uint32(i + 1),
// 			})
// 		}
// 	}
// 	byts, _ := json.Marshal(save)
// 	os.WriteFile("cn1.txt", byts, 0644)
// 	// buffranklist := bytes.NewBuffer([]byte{})
// 	// for _, rl := range dat {
// 	// 	for i, cell := range rl.Rank {
// 	// 		buffranklist.WriteString(fmt.Sprintf("%d\t%d\t%d\t%d\t%d\t%s\n", rl.GroupID, i+1, cell.Id>>16, cell.Id, cell.Score, cell.UnionName))
// 	// 	}
// 	// }
// 	// os.WriteFile("ranklist.txt", buffranklist.Bytes(), 0644)
// 	return nil
// }

// func loadWorld(cfg *pb.Jmongo, name string, rs map[uint32]uint32) {
// 	cfg.Db = name
// 	cli := xmongo.NewMongoClient(cfg)
// 	dat := &ActInfoDB{}
// 	cli.Exec(func(db *mongo.Database) error {
// 		err := db.Collection("act_bak").FindOne(context.TODO(), bson.M{"actid": 23101601}).Decode(dat)
// 		return err
// 	})

// 	for _, v := range dat.Data.Rank.UnionTaskDrawRank.Data {
// 		rs[uint32(v.Key)] = 0
// 	}

// 	buffranklist := bytes.NewBuffer([]byte{})
// 	for rank, u := range rs {
// 		if u == 0 {
// 			continue
// 		}
// 		cli.Exec(func(db *mongo.Database) error {
// 			uData := &pb.MsgUnion{}
// 			err := db.Collection("union").FindOne(context.TODO(), bson.M{"uid": u}).Decode(uData)
// 			for id, m := range uData.Member {
// 				buffranklist.WriteString(fmt.Sprintf("%d\t%d\t%d\t%s\n", rank, id, u, m.Name))
// 			}
// 			return err
// 		})
// 	}
// 	cli.Disconnect()
// 	os.WriteFile(name+".txt", buffranklist.Bytes(), 0644)
// }

func main() {
	cfg := &pb.MsgSysConfigs{}
	xloadcfg.Run("../../configs", cfg, nil)
	xlog.Shared(cfg.Logger, "examples")

	// loadRLs(cfg.Mongo)
	// //loadWorld(cfg.Mongo, "yw_czs_game9", rs[9])
	// // for sid, val := range rs {
	// // 	loadWorld(cfg.Mongo, fmt.Sprintf("gf_game%d", sid), val)
	// // }

	//创建索引
	//index()
	//插入数据
	//save()
	//读取数据
	//load()
}

func index() {
	ids := map[string]mongo.IndexModel{}
	ids["firstname_1_lastname_-1"] = mongo.IndexModel{
		Keys:    bson.D{{"firstname", 1}, {"lastname", -1}},
		Options: options.Index().SetUnique(false),
	}
	err := xmongo.CreateIndex("table", ids)
	if err != nil {
		xlog.Warn(err.Error())
	}

	// xmongo.ExecSync(func(db *mongo.Database) error {
	// 	//
	// 	model := mongo.IndexModel{
	// 		Keys:    bson.D{{"firstName", 1}, {"lastName", -1}},
	// 		Options: options.Index().SetUnique(false),
	// 	}
	// 	name, err := db.Collection("table").Indexes().CreateOne(context.TODO(), model)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	xlog.Infof("collection table create index %s", name)
	// 	return nil
	// })
}

func save() {
	err := xmongo.ExecSync(func(db *mongo.Database) error {
		dat := &Data{
			Firstname: "fname",
			Lastname:  "lname",
			Others:    "ots",
		}
		filter := bson.M{"firstname": "fname", "lastname": "lname"}
		opt := options.Replace().SetUpsert(true)
		_, err := db.Collection("table").ReplaceOne(context.TODO(), filter, dat, opt)
		return err
	})
	if err != nil {
		xlog.Warn(err.Error())
	}
}

func load() {
	dat := &Data{}
	err := xmongo.ExecSync(func(db *mongo.Database) error {
		i, err := db.Collection("table").CountDocuments(context.TODO(), bson.M{})
		if err == nil {
			xlog.Info(i)
		}
		err = db.Collection("table").FindOne(context.TODO(), bson.D{{"firstname", "fname"}, {"lastname", "lname"}}).Decode(dat)
		return err
	})
	if err != nil {
		xlog.Warn(err.Error())
	}
	xlog.Info(dat)
}
