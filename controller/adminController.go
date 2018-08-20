package controller

import(
	"github.com/gin-gonic/gin"
	"mtime.com/corednsUI/dao/etcd"
	"mtime.com/corednsUI/model"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/websocket"
	"fmt"
	"time"
)

func Oninit(){
	go etcd.WatchEtcd()
}

func DnsEditGet() func(*gin.Context)  {
	return func (c *gin.Context){
		key:= c.DefaultQuery("key","")
		if key == "" {
			c.HTML(http.StatusOK, "dnsedit.html",gin.H{})
		}else{
			obj,err:=etcd.EtcdDao.DnsGet(key)
			if err !=nil{
				c.HTML(http.StatusOK, "dnsedit.html",gin.H{"Message":model.Message{Error:err.Error()}})
			}else{
				c.HTML(http.StatusOK, "dnsedit.html",gin.H{"data":obj})
			}
		}
	}
}
func DnsEditPost() func(*gin.Context)  {
	return func (c *gin.Context){
		key :=c.DefaultQuery("key","")
		name:=c.PostForm("name")
		data:=c.PostForm("data")
		ttl:=c.PostForm("ttl")
		intTTL,_:= strconv.Atoi(ttl)
		if name =="" || data ==""{
			c.HTML(http.StatusOK, "dnsedit.html",gin.H{"obj":&model.Dns{Origin:name,NameServer:data,TTL:intTTL,Key:key},"error":"名称和数据不能为空！"})
			return
		}
		var value []byte
		if intTTL == 0{
			value,_=json.Marshal(model.A{Host:data})
		}else{
			value,_=json.Marshal(model.A{Host:data,TTL:intTTL})
		}
		if key == ""{
			etcd.EtcdDao.DnsAdd(name,string(value))
		}else{
			etcd.EtcdDao.DnsEdit(key,string(value))
		}
		c.Redirect(http.StatusMovedPermanently,"/admin/dns")
		return
	}
}

func Dnslist() func(*gin.Context)  {
	return func (c *gin.Context){
		c.HTML(http.StatusOK, "dns2.html",gin.H{
		})
	}
}
func Dnslist2() func(*gin.Context)  {
	return func (c *gin.Context){
		c.HTML(http.StatusOK, "dns2.html",gin.H{
		})
	}
}
func WsHandler() func(context *gin.Context){
	return func(c *gin.Context) {
		var conn *websocket.Conn
		var err error
		Wsupgrader := websocket.Upgrader{
			ReadBufferSize:   1024,
			WriteBufferSize:  1024,
			HandshakeTimeout: 5 * time.Second,
			// 取消ws跨域校验
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err = Wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		go func(conn *websocket.Conn) {
			defer conn.Close()
			for {
				select{
				case message,ok := <- etcd.NewMessage:
					if !ok{
						conn.WriteMessage(websocket.CloseMessage,[]byte{})
					}
					conn.PingHandler()
					err :=conn.WriteJSON(message)
					if err !=nil{
						fmt.Println(err)
						etcd.NewMessage <- message
						break
					}
				}
			}
		}(conn)
	}
}

func DelDns() func(*gin.Context)  {
	return func (c *gin.Context){
		key :=c.Query("key")
		err :=etcd.EtcdDao.DnsDel(key)
		if err != nil{
			c.JSON(http.StatusOK,model.Message{Error:err.Error()})
			return
		}else{
			c.JSON(http.StatusOK,model.Message{Error:""})
			return
		}

	}
}

//api 接口
func DnsApiList() func(*gin.Context)  {
	return func (c *gin.Context){
		c.JSON(http.StatusOK,model.DnsData{Data:etcd.EtcdDao.DnsList()})
	}
}