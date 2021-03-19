// /*
//  * @Author: mengdaoshizhongxinyang
//  * @Date: 2021-03-19 13:56:11
//  * @Description:
//  */
package server

// import (
// 	"fmt"
// 	"net/http"
// 	"runtime/debug"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/gorilla/websocket"
// 	"github.com/prometheus/common/log"
// 	"github.com/tidwall/gjson"
// )

// type websocketServer struct {
// 	token          string
// 	eventConn      []*websocketConn
// 	eventConnMutex sync.Mutex
// 	handshake      string
// }
// type websocketClient struct {
// 	token string

// 	universalConn *websocketConn
// 	eventConn     *websocketConn
// }
// type websocketConn struct {
// 	*websocket.Conn
// 	sync.Mutex
// }

// var WebsocketServer = &websocketServer{}
// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func (s *websocketServer) Run(addr, authToken string) {
// 	s.token = authToken

// 	s.handshake = fmt.Sprintf(`{"_post_method":2,"meta_event_type":"lifecycle","post_type":"meta_event","sub_type":"connect","time":%d}`,
// 		time.Now().Unix())
// 	http.HandleFunc("/event", s.event)
// 	http.HandleFunc("/api", s.api)
// 	http.HandleFunc("/", s.any)
// 	go func() {
// 		log.Infof("CQ Websocket 服务器已启动: %v", addr)
// 		log.Fatal(http.ListenAndServe(addr, nil))
// 	}()
// }

// func (s *websocketServer) event(w http.ResponseWriter, r *http.Request) {
// 	if s.token != "" {
// 		if auth := r.URL.Query().Get("access_token"); auth != s.token {
// 			if auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2); len(auth) != 2 || auth[1] != s.token {
// 				log.Warnf("已拒绝 %v 的 Websocket 请求: Token鉴权失败", r.RemoteAddr)
// 				w.WriteHeader(401)
// 				return
// 			}
// 		}
// 	}
// 	c, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Warnf("处理 Websocket 请求时出现错误: %v", err)
// 		return
// 	}
// 	err = c.WriteMessage(websocket.TextMessage, []byte(s.handshake))
// 	if err != nil {
// 		log.Warnf("Websocket 握手时出现错误: %v", err)
// 		c.Close()
// 		return
// 	}

// 	log.Infof("接受 Websocket 连接: %v (/event)", r.RemoteAddr)

// 	conn := &websocketConn{Conn: c}

// 	s.eventConnMutex.Lock()
// 	s.eventConn = append(s.eventConn, conn)
// 	s.eventConnMutex.Unlock()
// }

// func (s *websocketServer) api(w http.ResponseWriter, r *http.Request) {
// 	if s.token != "" {
// 		if auth := r.URL.Query().Get("access_token"); auth != s.token {
// 			if auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2); len(auth) != 2 || auth[1] != s.token {
// 				log.Warnf("已拒绝 %v 的 Websocket 请求: Token鉴权失败", r.RemoteAddr)
// 				w.WriteHeader(401)
// 				return
// 			}
// 		}
// 	}
// 	c, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Warnf("处理 Websocket 请求时出现错误: %v", err)
// 		return
// 	}
// 	log.Infof("接受 Websocket 连接: %v (/api)", r.RemoteAddr)
// 	conn := &websocketConn{Conn: c}
// 	go s.listenApi(conn)
// }

// func (s *websocketServer) any(w http.ResponseWriter, r *http.Request) {
// 	if s.token != "" {
// 		if auth := r.URL.Query().Get("access_token"); auth != s.token {
// 			if auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2); len(auth) != 2 || auth[1] != s.token {
// 				log.Warnf("已拒绝 %v 的 Websocket 请求: Token鉴权失败", r.RemoteAddr)
// 				w.WriteHeader(401)
// 				return
// 			}
// 		}
// 	}
// 	c, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Warnf("处理 Websocket 请求时出现错误: %v", err)
// 		return
// 	}
// 	err = c.WriteMessage(websocket.TextMessage, []byte(s.handshake))
// 	if err != nil {
// 		log.Warnf("Websocket 握手时出现错误: %v", err)
// 		c.Close()
// 		return
// 	}
// 	log.Infof("接受 Websocket 连接: %v (/)", r.RemoteAddr)
// 	conn := &websocketConn{Conn: c}
// 	s.eventConn = append(s.eventConn, conn)
// 	s.listenApi(conn)
// }
// func (c *websocketClient) listenApi(conn *websocketConn, u bool) {
// 	defer conn.Close()
// 	for {
// 		_, buf, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Warnf("监听反向WS API时出现错误: %v", err)
// 			break
// 		}

// 		go conn.handleRequest(buf)
// 	}
// }
// func (s *websocketServer) listenApi(c *websocketConn) {
// 	defer c.Close()
// 	for {
// 		t, payload, err := c.ReadMessage()
// 		if err != nil {
// 			break
// 		}

// 		if t == websocket.TextMessage {
// 			go c.handleRequest(payload)
// 		}
// 	}
// }

// func (c *websocketConn) handleRequest(payload []byte) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			log.NewErrorLogger().Print("处置WS命令时发生无法恢复的异常：%v\n%s", err, debug.Stack())
// 			c.Close()
// 		}
// 	}()
// 	j := gjson.ParseBytes(payload)
// 	t := strings.ReplaceAll(j.Get("action").Str, "_async", "")
// 	log.Debugf("WS接收到API调用: %v 参数: %v", t, j.Get("params").Raw)
// 	if f, ok := wsApi[t]; ok {
// 		ret := f(j.Get("params"))
// 		if j.Get("echo").Exists() {
// 			ret["echo"] = j.Get("echo").Value()
// 		}
// 		c.Lock()
// 		defer c.Unlock()
// 		_ = c.WriteJSON(ret)
// 	} else {
// 		ret := coolq.Failed(1404, "API_NOT_FOUND", "API不存在")
// 		if j.Get("echo").Exists() {
// 			ret["echo"] = j.Get("echo").Value()
// 		}
// 		c.Lock()
// 		defer c.Unlock()
// 		_ = c.WriteJSON(ret)
// 	}
// }

// var wsApi = map[string]func(gjson.Result){}
