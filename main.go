/*
 * @Author: mengdaoshizhongxinyang
 * @Date: 2021-03-16 09:54:23
 * @Description:
 */
package main

import (
	"Chatin/global"
	"Chatin/server"
	"fmt"
)

func main() {
	global.Gorm()
	server.HttpServer.Run(fmt.Sprintf("%s:%d", "127.0.0.1", 8080), "", 6000)
	select {}
}
