package main

import (
	"fmt"
	"lark/com/dbtables"
	"lark/com/pkgs/xloadcfg"
	"lark/com/pkgs/xlog"
	"lark/com/pkgs/xmysql"
	"lark/com/utils"
	"lark/pb"
	"time"

	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

func main() {
	cfg := &pb.MsgSysConfigs{}
	xloadcfg.Run("../../configs", cfg, nil)
	xlog.Shared(cfg.Logger, "examples")

	xmysql.NewMysqlClient(cfg.Mysql)
	xmysql.AutoMigrate(&dbtables.Users{})

	//先添加数据
	add()
	//读取数据
	read()
	//更新操作
	update()
}

func add() {
	data := &pb.UserInfo{
		Id:      1,
		Name:    "name",
		Account: "account",
	}
	bts, _ := proto.Marshal(data)
	now := time.Now().Unix()
	for i := 1; i < 11; i++ {
		user := &dbtables.Users{
			Account:  utils.NewUUID(),
			Name:     fmt.Sprintf("lark%d", i),
			CreateAt: now,
			LoginAt:  now,
			Level:    uint32(i),
			Avatar:   uint32(i),
			Data:     bts,
		}
		xmysql.Create(user)
	}
}

func read() {
	user1 := &dbtables.Users{}
	xmysql.Select(user1, "name = ?", "lark1")
	fmt.Println(user1)
	msg := &pb.UserInfo{}
	proto.Unmarshal(user1.Data, msg)
	fmt.Println(msg)

	user2 := &dbtables.Users{}
	xmysql.SqlSelectOne(user2, "select * from users where id = ?", 2)
	fmt.Println(user2)

	// _, ret := xmysql.SqlSelectList([]*dbtables.Users{}, "select * from users where id > ?", 9)
	// for _, v := range ret.([]*dbtables.Users) {
	// 	fmt.Println(v)
	// }
}

func update() {
	xmysql.Update(&dbtables.Users{ID: 5}, &dbtables.Users{Name: "lark55", Avatar: 55, Level: 55})

	user1 := &dbtables.Users{ID: 5}
	xmysql.Select(user1)
	fmt.Println(user1)

	xmysql.Delete(&dbtables.Users{ID: 6})

	xmysql.SqlExec("update users set name=?,level=?,avatar=? where id=?", "lark77", 77, 77, 7)

	xmysql.Run(func(db *gorm.DB) {
		if db == nil {
			return
		}
		var users []*dbtables.Users
		db.Raw("select * from users where id > ?", 9).Scan(&users)
		for _, user := range users {
			fmt.Println(user)
		}
	})
}
