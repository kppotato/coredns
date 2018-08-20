package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

const (
	cookiename = "_u"
)

func Login_post() func(*gin.Context)  {
	return func (c *gin.Context){
		cookie := http.Cookie{Name:cookiename,Value:"ssdfsdf", MaxAge:50*60*60,Secure:false,HttpOnly:true,Domain:"localhost",Path:"localhost"}
		http.SetCookie(c.Writer,&cookie)
		c.Redirect(http.StatusMovedPermanently,"/admin/dns")
	}
}

func Login_get() func(*gin.Context)  {
	return func (c *gin.Context){
		cookie,err :=c.Request.Cookie(cookiename)
		if err ==nil{
			cookie.MaxAge=-1
			http.SetCookie(c.Writer,cookie)
		}
		fmt.Println("login--------------------")
		c.HTML(http.StatusOK, "login.html",gin.H{
			"title": "Posts",
		})
	}
}
func AuthRequired() gin.HandlerFunc{
	return func(c *gin.Context){
		//这一部分可以替换成从session/cookie中获取，
		_,err := c.Request.Cookie(cookiename)
		if err != nil{
			fmt.Print(err)
			c.Redirect(http.StatusMovedPermanently,"/login")
			return
		}
		c.Next()
		//username:=c.Query("username")
		//password:=c.Query("password")

		//if username=="ft" && password =="123"{
		//	c.JSON(http.StatusOK,gin.H{"message":"身份验证成功"})
		//	c.Next()  //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
		//}else {
		//	c.Abort()
		//	c.JSON(http.StatusUnauthorized,gin.H{"message":"身份验证失败"})
		//	return// return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
		//}
	}
}