package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//跨域验证
func CORSMiddleware(origin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// log.Println("RequestURL:", c.Request.RequestURI, "  Method:", method)

		// ctx.Header("Access-Control-Allow-Origin", "http://106.12.79.177:9995") //注意部署的时候进行替换
		// ctx.Header("Access-Control-Allow-Origin", "http://106.12.79.177:9995") //注意部署的时候进行替换
		// ctx.Header("Access-Control-Allow-Headers", "*")
		// ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		// c.Header("Access-Control-Expose-Headers", "Authorization")
		// ctx.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", origin)
		// c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, No-Cache, X-Requested-With, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With, Authorization")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()

		// if c.Request.Method == http.MethodOptions {
		// 	c.AbortWithStatus(200)
		// 	log.Println("Opetion")
		// }
		// c.Next()
		// log.Println("Next")

	}
}
