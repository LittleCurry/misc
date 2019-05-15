package globals

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
	"net/http"
	"github.com/LittleCurry/misc/err_msg"
	"github.com/LittleCurry/misc/driver"
)

func KeyAuth() gin.HandlerFunc {

	return func (c *gin.Context)  {
		token := c.DefaultQuery("token", "")
		if token == "" {
			token = c.Request.Header.Get("Authorization")
			if s := strings.Split(token, " "); len(s) == 2 {
				token = s[1]
			}
		}
		if len(token) <= 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err_msg.ErrMsg{"登录过期, 请重新登录"})
			return
		}

		userId, err1 := driver.RedisClient.Get("AccessToken" + "-" + token).Result()
		if err1 != nil {
			fmt.Println("err1: ", err1)
			c.AbortWithStatusJSON(http.StatusUnauthorized, err_msg.ErrMsg{"登录过期, 请重新登录"})
			return
		}
		c.Set("userId", userId)

		userType, err2 := driver.RedisClient.Get("AccessTokenType-" + token).Result()
		if err2 == nil {
			c.Set("userType", userType)
		} else {
			c.Set("userType", "moment_user")
		}
		c.Next()
	}
}

/**
 * 跨域
 */
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}


