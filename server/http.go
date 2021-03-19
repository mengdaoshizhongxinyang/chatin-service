/*
 * @Author: mengdaoshizhongxinyang
 * @Date: 2021-03-16 10:03:54
 * @Description:
 */
package server

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

var Console = make(chan os.Signal, 1)

type httpServer struct {
	engine *gin.Engine
	Http   *http.Server
}
type httpClient struct {
	secret  string
	addr    string
	timeout int32
}

var HttpServer = &httpServer{}

func (c *httpClient) Run(addr, secret string, timeout int32) {
	c.secret = secret
	c.addr = addr
	c.timeout = timeout
	if c.timeout < 5 {
		c.timeout = 5
	}
}
func (s *httpServer) Run(addr, secret string, timeout int32) {
	gin.SetMode((gin.ReleaseMode))
	s.engine = gin.New()
	s.engine.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	s.engine.Any("/:action", s.HandleActions)

	go func() {
		log.Infof("http serve 启动成功")
		s.Http = &http.Server{
			Addr:    addr,
			Handler: s.engine,
		}

		if err := s.Http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(err)
			log.Infof("HTTP 服务启动失败, 请检查端口是否被占用.")
			log.Warnf("将在五秒后退出.")
			time.Sleep(time.Second * 5)
			os.Exit(1)
		}

	}()

}
func NewHttpClient() *httpClient {
	return &httpClient{}
}

func (s *httpServer) HandleActions(c *gin.Context) {
	action := strings.ReplaceAll(c.Param("action"), "_async", "")
	if f, ok := httpApi[action]; ok {
		f(s, c)
	} else {
		println(s, ok)
		c.JSON(404, gin.H{
			"states": "404",
		})
	}
}
func getTestInfo(s *httpServer, c *gin.Context) {
	c.JSON(200, gin.H{
		"states":  "200",
		"message": "test",
	})
}
func login(s *httpServer, c *gin.Context) {

	c.JSON(200, &c)
}

var httpApi = map[string]func(s *httpServer, c *gin.Context){
	"test":  getTestInfo,
	"login": login,
}
