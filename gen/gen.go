package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"path/filepath"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

var genPath = "/src/github.com/frankxi/league/gen/Config.json"

// sql类型转换golang类型
var typeMap = map[string]string{
	"int":       "int",
	"decimal":   "float64",
	"varchar":   "string",
	"timestamp": "string",
	"datetime":  "string",
}

type Model struct {
	Comment string
	Entry   string
	Table   string
}

type DbInfo struct {
	Ip                 string
	Port               string
	Username           string
	Password           string
	Db                 string
	TableAndStructList string
	BizPath            string
}

type TableAndStructInfo struct {
	Table  string
	Struct string
}

type TableInfo struct {
	Field      string
	Type       string
	Collation  sql.NullString
	Null       string
	Key        string
	Default    sql.NullString
	Extra      string
	Privileges string
	Comment    string
}

func getCurrentDirectory() string {
	str, _ := os.Getwd()
	fmt.Println(str)
	dir, err := filepath.Abs(filepath.Dir(str))
	if err != nil {
		fmt.Errorf("%d", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func (db *DbInfo) ReadLocalFile() []byte {
	son := getCurrentDirectory()
	son += genPath
	file, err := ioutil.ReadFile(son)
	if err != nil {
		panic(err)
	}
	return file
}
func BuildConnInfo(db *DbInfo) (string) {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db.Username, db.Password, db.Ip, db.Port, db.Db)
}

func main() {
	//常量配置

	//dev5
	dbConfig := new(DbInfo)
	err := json.Unmarshal(dbConfig.ReadLocalFile(), dbConfig)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	url := BuildConnInfo(dbConfig)
	project := "common"
	//连接数据库
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
		return
	}
	defer db.Close()
	path := dbConfig.BizPath
	//项目配置
	comment := "系统" //实体类含义
	tableAndStructList := strings.Split(dbConfig.TableAndStructList, ",")
	tablelist := make(map[string]Model)
	for _, item := range tableAndStructList {
		tableAndStruct := strings.Split(item, "**")
		tablelist[tableAndStruct[1]] = Model{
			Table:   tableAndStruct[0],
			Entry:   tableAndStruct[1],
			Comment: tableAndStruct[2],
		}
	}
	tab := "	"
	PrimaryKey := "id"
	for entry, table := range tablelist {
		comment = table.Comment
		//创建目录
		fmt.Println("开始生成" + comment + "代码")
		os.MkdirAll(path+"models/", 0777)

		rows, err1 := db.Query("SHOW FULL FIELDS FROM " + table.Table)
		if err1 != nil {
			fmt.Println(err1.Error())
			return
		}
		defer rows.Close()
		var tis []TableInfo
		for rows.Next() {
			var ti TableInfo
			err2 := rows.Scan(&ti.Field, &ti.Type, &ti.Collation, &ti.Null, &ti.Key, &ti.Default, &ti.Extra, &ti.Privileges, &ti.Comment)
			if err2 != nil {
				fmt.Println(err2.Error())
				return
			}
			tis = append(tis, ti)
		}

		/***********************生成Entry***********************/
		var m bytes.Buffer
		m.WriteString("package models\n\n")
		m.WriteString("import (\n")
		m.WriteString("		\"github.com/jinzhu/gorm\"\n")
		m.WriteString("		\"github.com/frankxi/league/biz\"\n")
		m.WriteString("		\"errors\"\n")
		//m.WriteString("		\"time\"\n\n")
		m.WriteString("\n) \n")
		m.WriteString("// " + comment + "类\n")
		m.WriteString("type " + entry + " struct {\n")
		//生成字段
		//m.WriteString("//   " + entry + " struct {\n")
		var jsonBf bytes.Buffer
		var columnBf bytes.Buffer
		jsonBf.WriteString("{")
		columnBf.WriteString("")
		for _, v := range tis {
			columnBf.WriteString(v.Field + ",")
			if v.Key == "PRI" {
				PrimaryKey = TtoF(v.Field)
				m.WriteString(tab + TtoF(v.Field) + tab + "int64" + tab + "`json:\"" + v.Field + "\" " + "gorm:\"primary_key\"`" + tab + "// id\n")
				jsonBf.WriteString("\"" + v.Field + "\":" + "0,")
			} else {
				var def = v.Default
				var t string
				for key, value := range typeMap {
					if strings.Contains(v.Type, key) {
						t = value
						break
					}
				}
				if t == "" {
					t = "string"
				}
				//默认值
				defaultStr := ""
				if def.Valid {
					defaultStr = "gorm:\"default:'" + def.String + "'\""
				}

				if t == "Time" {

					//m.WriteString(tab + TtoF(v.Field) + tab + "time." + t + tab + "`json:\"" + Camel(v.Field) + "\"` " + tab + "// " + v.Comment + "\n")
					m.WriteString(tab + TtoF(v.Field) + tab + "string" + tab + "`json:\"" + Camel(v.Field) + "\" " + defaultStr + "`" + tab + "// " + v.Comment + "\n")
				} else {
					m.WriteString(tab + TtoF(v.Field) + tab + t + tab + "`json:\"" + Camel(v.Field) + "\" " + defaultStr + "`" + tab + "// " + v.Comment + "\n")
				}
				if t == "int" {
					jsonBf.WriteString("\"" + Camel(v.Field) + "\":" + "0,")
				} else {
					jsonBf.WriteString("\"" + Camel(v.Field) + "\":" + "\"\",")
				}

			}
		}
		m.WriteString("}")
		columnJson := columnBf.String()[0 : len(columnBf.String())-1]
		m.WriteString("\n//" + columnJson)
		json := jsonBf.String()[0:len(jsonBf.String())-1] + "}"
		fmt.Println(columnJson)
		fmt.Println(path + "models/model" + entry + "_CRUD.go")
		ioutil.WriteFile(path+"models/model"+entry+"_CRUD.go", []byte(m.String()), 0777)
		fmt.Println("Entry生成完毕")

		/***********************生成dao***********************/
		var dao = `

// 分页查询条件
type {model}Query struct {
	{model}
	biz.Page
}

// 分页查询结果
type {model}Page struct {
	List []{model} {sep}json:"list"{sep}
	biz.Page
}

// 主键查询(注意gorm.ErrRecordNotFound)
func (m *{model}) GetById(db *gorm.DB, id int) ({model}, error) {
	var one {model}
	return one, db.Model(m).First(&one, id).Error
}

// 条件查询
func (m *{model}) QueryList(db *gorm.DB, filter map[string]interface{}) ([]{model}, error) {
	var result []{model}
	return result, db.Model(*m).Where(filter).Find(&result).Error
}

// 分页查询
func (m *{model}Query) QueryPage(db *gorm.DB) ({model}Page, error) {
	var page {model}Page
	page.Page = m.Page
	db = db.Model(m.{model}).Where(m.{model})
	err := db.Count(&page.Total).Error
	if err != nil {
		return page, err
	}
	err = db.Offset((page.PageNo - 1) * page.PageSize).Limit(page.PageSize).Find(&page.List).Error
	return page, err
}

// 新增
func (m *{model}) Insert(db *gorm.DB) error {
	return db.Model(m).Create(m).Error
}

// 更新(注意不能自动更新值为''或0)
func (m *{model}) Update(db *gorm.DB) error {
	if m.{PrimaryKey} == 0 {
		return errors.New("没有查询条件！")
	}
	return db.Model(m).Update(m).Error
}

// 更新
func (m *{model}) UpdateWhere(db *gorm.DB, filter map[string]interface{}, fields map[string]interface{}) error {
	// 检查查询条件
	if len(filter) == 0 {
		return errors.New("没有查询条件！")
	}
	return db.Model(m).Where(filter).Update(fields).Error
}
`
		dao = strings.Replace(dao, "{comment}", comment, -1)
		dao = strings.Replace(dao, "{project}", project, -1)
		dao = strings.Replace(dao, "{model}", entry, -1)
		dao = strings.Replace(dao, "{PrimaryKey}", PrimaryKey, -1)
		dao = strings.Replace(dao, "{sep}", "`", -1)

		f, err := os.OpenFile(path+"models/model"+entry+"_CRUD.go", os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
		} else {
			// 查找文件末尾的偏移量
			n, _ := f.Seek(0, 2)
			// 从末尾的偏移量开始写入内容
			_, err = f.WriteAt([]byte(dao), n)
		}
		defer f.Close()
		fmt.Println("Dao生成完毕")
		/***********************生成service***********************/
		var service = `package process

// {comment}业务类
import (
	"../models"
	"github.com/frankxi/league/dao"
)

// ID查询
func Get{model}(id int) (models.{model}, error) {
	var q models.{model}
	return q.GetById(dao.GetDB(), id)
}

// 列表
func List{model}(query map[string]interface{}) ([]models.{model}, error) {
	var q models.{model}
	return q.QueryList(dao.GetDB(), query)
}

// 分页
func Page{model}(query models.{model}Query) (models.{model}Page, error) {
	return query.QueryPage(dao.GetDB())
}

// 新增
func Insert{model}(m *models.{model}) error {
	return m.Insert(dao.GetDB())
}

// 更新
func Update{model}(m *models.{model}) error {
	return m.Update(dao.GetDB())
}

// 删除
func Delete{model}(m *models.{model}) error {
	return nil
}
`
		service = strings.Replace(service, "{comment}", comment, -1)
		service = strings.Replace(service, "{project}", project, -1)
		service = strings.Replace(service, "{model}", entry, -1)
		//fmt.Println(service)
		ioutil.WriteFile(path+"process/process"+entry+".go", []byte(service), 0777)
		fmt.Println("Service生成完毕")
		/***********************生成api***********************/
		var api = `package api

import (
	"github.com/gin-gonic/gin"
     log "github.com/sirupsen/logrus"
	"../models"
	"../process"
	"github.com/frankxi/league/comm/e"
	"github.com/frankxi/league/utils"
)

// 新增{comment}
func insert{model}(c *gin.Context) {
	var p models.{model}
	if err := c.Bind(&p); err != nil {
		log.Error("insert{model}=", p)
        utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("insert{model}=", p)
	err := process.Insert{model}(&p)
	if err != nil {
		log.Error("insert{model}.error: %+v", err.Error())
        utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}

// 更新{comment}
func update{model}(c *gin.Context) {
	var p models.{model}
	if err := c.Bind(&p); err != nil {
		log.Error("update{model}=", p)
        utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("update{model}=", p)
	err := process.Update{model}(&p)
	if err != nil {
		log.Error("update{model}.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}

// 删除{comment}
func delete{model}(c *gin.Context) {
	var p models.{model}
	if err := c.Bind(&p); err != nil {
		log.Error("delete{model}=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("delete{model}=", p)
	err := process.Delete{model}(&p)
	if err != nil {
		log.Error("delete{model}.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}
`
		api = strings.Replace(api, "{comment}", comment, -1)
		api = strings.Replace(api, "{project}", project, -1)
		api = strings.Replace(api, "{smallModel}", Camel(entry), -1)
		api = strings.Replace(api, "{model}", entry, -1)
		//api = strings.Replace(api, "{json}", json, -1)
		//fmt.Println(service)
		ioutil.WriteFile(path+"api/api"+entry+".go", []byte(api), 0777)
		fmt.Println("Api生成完毕")

		/***********************生成console***********************/
		//c.JSON(ERR, "参数错误！")
		//c.JSON(ERR, err.Error())
		//c.JSON(OK, nil)
		//log.Error
		//log.Info
		var console = `package console

import (
	"github.com/gin-gonic/gin"
	"../models"
	"../process"
 	log "github.com/sirupsen/logrus"
	"github.com/frankxi/league/comm/e"
	"github.com/frankxi/league/utils"
)

// 注册API{comment}
// {json}
// 所有字段
// {columnJson}
func AddConsole{model}(base *gin.RouterGroup) {
	api := base.Group("/{smallModel}")
	api.POST("/list", list{model})
	api.POST("/insert", insert{model})
	api.POST("/update", update{model})
	api.POST("/delete", delete{model})
}
// 分页获取{comment}
func list{model}(c *gin.Context) {
	var p models.{model}Query
	if err := c.Bind(&p); err != nil {
		log.Error("list{model}=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("list{model}=", p)
	page, err := process.Page{model}(p)
	if err != nil {
		log.Error("list{model}.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,page)
}

// 新增{comment}
func insert{model}(c *gin.Context) {
	var p models.{model}
	if err := c.Bind(&p); err != nil {
		log.Error("insert{model}=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("insert{model}=", p)
	err := process.Insert{model}(&p)
	if err != nil {
		log.Error("insert{model}.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}

// 更新{comment}
func update{model}(c *gin.Context) {
	var p models.{model}
	if err := c.Bind(&p); err != nil {
		log.Error("update{model}=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("update{model}=", p)
	err := process.Update{model}(&p)
	if err != nil {
		log.Error("update{model}.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}

// 删除{comment}
func delete{model}(c *gin.Context) {
	var p models.{model}
	if err := c.Bind(&p); err != nil {
		log.Error("delete{model}=", p)
		utils.Err(c, e.ErrParamCode)
		return
	}
	log.Info("delete{model}=", p)
	err := process.Delete{model}(&p)
	if err != nil {
		log.Error("delete{model}.error: %+v", err.Error())
		utils.Err(c, e.ErrSystemCode)
		return
	}
	utils.OK(c,nil)
}
`
		console = strings.Replace(console, "{comment}", comment, -1)
		console = strings.Replace(console, "{project}", project, -1)
		console = strings.Replace(console, "{smallModel}", Camel(entry), -1)
		console = strings.Replace(console, "{model}", entry, -1)
		console = strings.Replace(console, "{json}", json, -1)
		console = strings.Replace(console, "{columnJson}", columnJson, -1)
		//fmt.Println(service)
		ioutil.WriteFile(path+"console/console"+entry+".go", []byte(console), 0777)
		fmt.Println("console生成完毕")
	}
}

func TtoF(table string) string {
	source := strings.Split(table, "_")
	ret := ""
	for _, v := range source {
		ret = ret + strings.ToUpper(v[0:1]) + v[1:]
	}
	return ret
}

func Camel(table string) string {
	res := TtoF(table)
	res = strings.ToLower(res[0:1]) + res[1:]
	return res
}
