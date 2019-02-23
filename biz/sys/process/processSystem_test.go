package process

import (
	"testing"
	"fmt"
	"git.woda.ink/jifanfei/common/utils/DbGorm"
	"gitee.com/league/dao"
	"gitee.com/league/biz/sys/models"
)

func init() {

	//init gorm
	gGorm, err := dao.InitDb()
	if err != nil {
		return
	}
	//设置DB
	dao.SetDB(gGorm)
	// 默认打印所有gorm的sql日志
	gGorm.LogMode(true)
	//打印gorm的sql
	gGorm.SetLogger(new(DbGorm.QsLogImpl))
}

func TestGetSystem(t *testing.T) {
	//Init()
	res, err := GetSystem(1)
	if err != nil {
		//t.Errorf("err",err)
		t.Fatal(err)

	}
	fmt.Println(res)
}

func TestListSystem(t *testing.T) {
	//Init()
	var p models.System
	p.Name = "1"
	res, err := ListSystem(p)
	if err != nil {
		//t.Errorf("err",err)
		t.Fatal(err)

	}
	fmt.Println(res)
}
