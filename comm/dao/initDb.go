package dao

import (
	"github.com/jinzhu/gorm"
	"fmt"
	log "github.com/sirupsen/logrus"
	"errors"
	"time"
	"sync"
	"github.com/frankxi/league/comm/setting"
	"gitee.com/league_bak/dao"
	"git.woda.ink/jifanfei/common/utils/DbGorm"
)

var dbPingOnce sync.Once
var gGorm *gorm.DB

func Setup() {
	//init gorm
	gGorm, err := InitDb()
	if err != nil {
		log.Error("InitDb", "gorm.Open", err)
		return
	}
	//设置DB
	dao.SetDB(gGorm)
	// 默认打印所有gorm的sql日志
	gGorm.LogMode(true)
	//打印gorm的sql
	gGorm.SetLogger(new(DbGorm.QsLogImpl))
}

func GetDB() *gorm.DB {
	return gGorm
}

func SetDB(db *gorm.DB) {
	gGorm = db
}

func InitDb() (*gorm.DB, error) {
	ggorm, err := gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Error("InitDb", "gorm.Open", err)
		return nil, errors.New("gorm.Open")
	}
	err = ggorm.DB().Ping()
	if err != nil {
		log.Error("gORM.DB().Ping", "gorm.Open", err)
		return nil, errors.New("gORM.DB().Ping")
	}
	// 连接池属性设置
	ggorm.DB().SetMaxIdleConns(setting.DatabaseSetting.MaxIdleConns)
	ggorm.DB().SetMaxOpenConns(setting.DatabaseSetting.MaxOpenConns)

	dbPingOnce.Do(func() {
		go dbTimerPing(ggorm)
	})

	// 全局禁用表名复数
	// 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	ggorm.SingularTable(true)

	return ggorm, nil
}

/*
 * @brief 创建mysql数据库链接信息
 *
 * @param user 用户名
 * @param pwd 密码
 * @param ip 数据库地址
 * @param port 端口
 * @param db 数据库
 * @return 数据库链接信息
 */
func BuildConnInfo(user string, pwd string, ip string, port string, db string) (string) {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pwd, ip, port, db)
}

func dbTimerPing(ggorm *gorm.DB) {
	tick := time.NewTicker(time.Second * 300)

	for {
		select {
		case <-tick.C:
			ggorm.DB().Ping()
		}
	}
}
