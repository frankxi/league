package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/frankxi/league/comm/setting"
	"github.com/frankxi/league/comm/dao"
	"github.com/frankxi/league/comm/l"
	"github.com/frankxi/league/comm/gredis"
	"github.com/frankxi/league/routers"
	"net/http"
	"fmt"
	"runtime/debug"
	_ "./docs" // docs is generated by Swag CLI, you have to import it.
)

var server *http.Server

func init() {
	setting.Setup()    //setup  config
	dao.Setup()        //setup db connection
	logSetting.Setup() //setup log formatter
	gredis.Setup()     //setup redis tcp connection
	//init gin & routers
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	//config http service
	server = &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", server.Addr)
}
// @title League
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://coding13.com/terms/

// @contact.name API Support
// @contact.url http://coding13.com/support
// @contact.email ogavaj@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host coding13.com
// @BasePath /v1
func main() {
	//crash print stack
	defer func() {
		if err := recover(); err != nil {
			log.Error("main", "server crash: ", err)
			log.Error("main", "stack: ", string(debug.Stack()))
		}
	}()

	//listen and serve
	err := server.ListenAndServe()
	if err != nil {
		log.Error("ListenAndServe ", "error", err)
	}
}
