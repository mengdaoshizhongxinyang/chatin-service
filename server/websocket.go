/*
 * @Author: mengdaoshizhongxinyang
 * @Date: 2021-03-19 13:56:11
 * @Description:
 */
package server

import (
	"Chatin/global"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var (
	clientManager = NewClientManager()
	appIds []uint32

	serverIp string
	serverPort string
)

func GetAppIds() []uint32 {
	return appIds
}

func StartWebSocket(){

	serverIp=global.GetServerIp()
	webSocketPort :=viper.GetString("app.webSocketPort")
	rpcPort:=viper.GetString("app.rpcPort")

	serverPort=rpcPort
	http.HandlerFunc("/",wsPage)

	go clientManager.start()
}

func wsPage(w http.ResponseWriter, req *http.Request) {

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	clientManager.Register <- client
}