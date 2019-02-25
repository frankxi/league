package dao

import (
	"github.com/jinzhu/gorm"
	"fmt"
	log "github.com/sirupsen/logrus"
	"errors"
	"time"
	"sync"
	"github.com/frankxi/league/comm/setting"
	"reflect"
	"database/sql/driver"
	"regexp"
	"strconv"
)

var dbPingOnce sync.Once
var gGorm *gorm.DB
var (
	sqlRegexp                = regexp.MustCompile(`\?`)
	numericPlaceHolderRegexp = regexp.MustCompile(`\$\d+`)
)

func Setup() {
	//init gorm
	gGorm, err := InitDb()
	if err != nil {
		log.Error("InitDb", "gorm.Open", err)
		return
	}
	//设置DB
	SetDB(gGorm)
	// 默认打印所有gorm的sql日志
	gGorm.LogMode(true)
	//打印gorm的sql
	gGorm.SetLogger(new(QsLogImpl))
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

type LogImpl struct{}

func (this *LogImpl) Print(args ...interface{}) {
	str := ""
	for i, v := range args {
		if i == 0 {
			str = str + fmt.Sprint(v)
		} else {
			str = str + fmt.Sprint(" ", v)
		}
	}
	log.Info("gorm", "data", str)
}

var NowFunc = func() time.Time {
	return time.Now()
}

var LogFormatter = func(values ...interface{}) (messages []interface{}) {
	if len(values) > 1 {
		var (
			sql             string
			formattedValues []string
			level           = values[0]
			currentTime     = "\n\033[33m[" + NowFunc().Format("2006-01-02 15:04:05") + "]\033[0m"
			source          = fmt.Sprintf("\033[35m(%v)\033[0m", values[1])
		)

		messages = []interface{}{source, currentTime}

		if level == "sql" {
			// duration
			messages = append(messages, fmt.Sprintf(" \033[36;1m[%.2fms]\033[0m ", float64(values[2].(time.Duration).Nanoseconds()/1e4)/100.0))
			// sql

			for _, value := range values[4].([]interface{}) {
				indirectValue := reflect.Indirect(reflect.ValueOf(value))
				if indirectValue.IsValid() {
					value = indirectValue.Interface()
					if t, ok := value.(time.Time); ok {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", t.Format("2006-01-02 15:04:05")))
					} else if b, ok := value.([]byte); ok {
						if str := string(b); true {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", str))
						} else {
							formattedValues = append(formattedValues, "'<binary>'")
						}
					} else if r, ok := value.(driver.Valuer); ok {
						if value, err := r.Value(); err == nil && value != nil {
							formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
						} else {
							formattedValues = append(formattedValues, "NULL")
						}
					} else {
						formattedValues = append(formattedValues, fmt.Sprintf("'%v'", value))
					}
				} else {
					formattedValues = append(formattedValues, "NULL")
				}
			}

			// differentiate between $n placeholders or else treat like ?
			if numericPlaceHolderRegexp.MatchString(values[3].(string)) {
				sql = values[3].(string)
				for index, value := range formattedValues {
					placeholder := fmt.Sprintf(`\$%d([^\d]|$)`, index+1)
					sql = regexp.MustCompile(placeholder).ReplaceAllString(sql, value+"$1")
				}
			} else {
				formattedValuesLength := len(formattedValues)
				for index, value := range sqlRegexp.Split(values[3].(string), -1) {
					sql += value
					if index < formattedValuesLength {
						sql += formattedValues[index]
					}
				}
			}

			messages = append(messages, sql)
			messages = append(messages, fmt.Sprintf(" \n\033[36;31m[%v]\033[0m ", strconv.FormatInt(values[5].(int64), 10)+" rows affected or returned "))
		} else {
			messages = append(messages, "\033[31;1m")
			messages = append(messages, values[2:]...)
			messages = append(messages, "\033[0m")
		}
	}

	return
}

type QsLogImpl struct{}

func (this *QsLogImpl) Print(args ...interface{}) {
	messages := LogFormatter(args...)
	str := fmt.Sprintln(messages)
	log.Info("gorm", "data", str)
}
