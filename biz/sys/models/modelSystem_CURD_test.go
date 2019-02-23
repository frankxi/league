package models

import (
	"testing"
	"fmt"
	"git.woda.ink/jifanfei/common/utils/DbGorm"
	"gitee.com/league/dao"
)

var system = new(System)

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
	res, err := system.GetById(dao.GetDB(), 1)
	if err != nil {
		//t.Errorf("err",err)
		t.Fatal(err)

	}
	fmt.Println(res)
}

func TestListSystem(t *testing.T) {
	//Init()
	params := make(map[string]interface{}, 0)
	//params["Name"] ="1"
	res, err := system.QueryList(dao.GetDB(), params)
	if err != nil {
		//t.Errorf("err",err)
		t.Fatal(err)

	}
	fmt.Println(res)
}

func TestInsert(t *testing.T) {
	//Init()
	system = &System{
		Name: "3",
		Code: "1",
		Desc: "3",
		Logo: "3",
		IsEnabled:10,
	}
	err := system.Insert(dao.GetDB())
	if err != nil {
		//t.Errorf("err",err)
		t.Fatal(err)

	}
}

func TestUpdate(t *testing.T) {
	//Init()
	system = &System{
		Id:   2,
		Name: "3",
		Code: "-4",
		Desc: "xxxx3",
		Logo: "3",
	}
	err := system.Update(dao.GetDB())
	if err != nil {
		//t.Errorf("err",err)
		t.Fatal(err)

	}
}


func TestUpdateWhere (t *testing.T) {
	where  := make(map[string]interface{}, 0)
	where["Id"] = 2
	where["Name"] = "3"

	fileds  := make(map[string]interface{}, 0)
	fileds["Code"] = "-xxxx"
	fileds["Desc"] = "-yyyy"
	err :=system.UpdateWhere(dao.GetDB(),where,fileds)
	if err != nil {
		//t.Errorf("err",err)
		t.Fatal(err)

	}
}