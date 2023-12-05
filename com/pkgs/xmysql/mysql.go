package xmysql

import (
	"database/sql"
	"errors"
	"fmt"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"
	"reflect"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlClient struct {
	db  *gorm.DB
	cfg *pb.Jmysql
}

var cli *MysqlClient

func NewMysqlClient(cfg *pb.Jmysql) *MysqlClient {
	if cli == nil {
		cli = &MysqlClient{
			cfg: cfg,
		}
		cli.connectDB()
		return cli
	} else {
		cli.cfg = cfg
		return cli
	}
}

// 新建连接
func (cli *MysqlClient) connectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cli.cfg.Username,
		cli.cfg.Password,
		cli.cfg.Address,
		"",
	)
	opts := &gorm.Config{
		SkipDefaultTransaction: false, // 禁用默认事务(true: Error 1295: This command is not supported in the prepared statement protocol yet)
		PrepareStmt:            false, // 创建并缓存预编译语句(true: Error 1295)
	}
	db, err := gorm.Open(mysql.Open(dsn), opts)
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	cli.db = db
	cli.usedb()
}

func (cli *MysqlClient) usedb() {
	err := cli.db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci ", cli.cfg.Db)).Error
	if err != nil {
		xlog.Error(err.Error())
		return
	}

	err = cli.db.Exec(fmt.Sprintf("use %s", cli.cfg.Db)).Error
	if err != nil {
		xlog.Error(err.Error())
		return
	}

	sqlDB, err := cli.db.DB()
	if err != nil {
		xlog.Error(err.Error())
		return
	}
	//设置最大空闲连接
	sqlDB.SetMaxIdleConns(int(cli.cfg.MaxIdleConn))
	//设置最大连接数
	sqlDB.SetMaxOpenConns(int(cli.cfg.MaxOpenConn))
	//设置连接超时时间
	sqlDB.SetConnMaxLifetime(time.Duration(cli.cfg.ConnLifetime) * time.Second)
}

func isConn() bool {
	if cli == nil || cli.db == nil {
		xlog.Error("please connect db first")
		return false
	}
	return true
}

// 同步数据表
func AutoMigrate(tables ...interface{}) {
	if !isConn() {
		return
	}
	cli.db.AutoMigrate(tables...)
}

/*
	新建一条数据

Create(&User{...})
*/
func Create(table interface{}) bool {
	if !isConn() {
		return false
	}
	result := cli.db.Create(table)
	if result.Error != nil {
		xlog.Error(result.Error.Error())
		return false
	}
	return true
}

/*
	获取一条记录

Select(&User{ID:1})
Select(&User{...}, "name=?", "xxx")
*/
func Select(table interface{}, condations ...interface{}) bool {
	if !isConn() {
		return false
	}
	result := cli.db.First(table, condations...)
	if result.Error != nil {
		xlog.Error(result.Error.Error())
		return false
	}
	return true
}

/*
	更新一条数据

Update(&User{ID:1}, &User{Name:"update", ...})
*/
func Update(key, table interface{}) bool {
	if !isConn() {
		return false
	}
	result := cli.db.Model(key).Updates(table)
	if result.Error != nil {
		xlog.Error(result.Error.Error())
		return false
	}
	return true
}

/*
	删除一条数据

Delete(&User{ID:1})
*/
func Delete(table interface{}) bool {
	if !isConn() {
		return false
	}
	result := cli.db.Delete(table)
	if result.Error != nil {
		xlog.Error(result.Error.Error())
		return false
	}
	return true
}

/*
	单条查询

SqlSelectOne(&dbtables.Users{}, "select * from users where id = ?", 1)
*/
func SqlSelectOne(table interface{}, sql string, params ...interface{}) bool {
	if !isConn() {
		return false
	}
	result := cli.db.Raw(sql, params...).Scan(table)
	if result.Error != nil {
		xlog.Error(result.Error.Error())
		return false
	}
	return true
}

/*
	列表查询

b, ret := SqlSelectList([]*User{}, "select * from users where id > ?", 9)
结果: ret.([]*User)
*/
func SqlSelectList(table interface{}, sql string, params ...interface{}) (bool, interface{}) {
	if !isConn() {
		return false, nil
	}
	rows, err := cli.db.Raw(sql, params...).Rows()
	if err != nil {
		xlog.Error(err.Error())
		return false, nil
	}
	defer rows.Close()

	arr := reflect.ValueOf(table)
	typ := reflect.TypeOf(table).Elem().Elem()
	for rows.Next() {
		dat := reflect.New(typ)
		cli.db.ScanRows(rows, dat.Interface())
		arr = reflect.Append(arr, dat)
	}
	return true, arr.Interface()
}

/*
	sql更新操作

SqlExec("update users set name=?,level=?,avatar=? where id=?", "lark77", 77, 77, 7)
*/
func SqlExec(sql string, params ...interface{}) bool {
	if !isConn() {
		return false
	}
	result := cli.db.Exec(sql, params...)
	if result.Error != nil {
		xlog.Error(result.Error.Error())
		return false
	}
	return true
}

/*
自定义操作

	func(db *gorm.DB) {
		if db == nil {
			return
		}
		var users []*User
		db.Raw("select * from users where id > ?", 9).Scan(&users)
		for _, user := range users {
			fmt.Println(user)
		}
	}
*/
func Run(f func(*gorm.DB)) {
	f(cli.db)
}

// 事务处理
func Transaction(handle func(*gorm.DB) error) error {
	if cli == nil {
		return errors.New(utils.Const_Mysql_Not_Instance)
	}
	if cli.db == nil {
		return errors.New(utils.Const_Mysql_Not_Connect)
	}

	tx := cli.db.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	err := handle(tx)
	if err != nil {
		xlog.Error(err.Error())
		return tx.Rollback().Error
	}
	return tx.Commit().Error
}
