package httpGin

import (
	"github.com/gin-gonic/gin"
	"github.com/kppotato/coredns/controller"
	"net/http"
)

func StartHttp() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.Static("/static", "D:\\project\\go\\src\\mtime.com\\corednsUI\\static")
	admin := r.Group("admin")
	admin.Use(controller.AuthRequired())
	{
		admin.GET("", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/admin/dns")
		})
		admin.GET("dns", controller.Dnslist2())
		//admin.GET("dns2",controller.Dnslist2())
		admin.GET("deldns", controller.DelDns())
		admin.GET("etcdlist", controller.DnsApiList())
		admin.GET("dnsedit", controller.DnsEditGet())
		admin.POST("dnsedit", controller.DnsEditPost())
		//admin.GET("ws",controller.WsHandler())
	}
	r.GET("ws", controller.WsHandler())
	r.GET("/login", controller.Login_get())
	r.POST("/login", controller.Login_post())
	r.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin/dns")
	})
	r.Run(":9191")
}
