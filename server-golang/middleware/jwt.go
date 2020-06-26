package middleware

import (
	"errors"
	"labelproject-back/common"
	"labelproject-back/model"
	"labelproject-back/util"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"time"
)

var jwtKey = []byte("a_secret_crect")

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type Claims struct {
	UserName string
	jwt.StandardClaims
	IP map[string]string
}

// ReleaseToken 发放Token
func ReleaseToken(ctx *gin.Context, user model.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &Claims{
		UserName: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		IP: map[string]string{"ip": ctx.Request.RemoteAddr},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return "", TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return "", TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return "", TokenNotValidYet
			} else {
				return "", TokenInvalid
			}
		}
	}

	return tokenString, nil

}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}

// TokenIsExpiration is to compare the
func TokenIsExpiration(expirationTime string) (bool, error) {
	// log.Println("过期时间：", expirationTime, "\n")
	Oldtime, _ := time.ParseInLocation("2006-01-02 15:04:05", expirationTime[:19], time.Local)
	// log.Println("解析后时间： ", Oldtime, "\n")
	now := time.Now().Local()
	return now.After(Oldtime), nil
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookie, _ := c.Request.Cookie("request_token")
		if cookie == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "请先登陆"})
			c.Abort()
		}

		authorization := c.Request.Header.Get("Authorization")
		// var username string = ""
		// log.Println("Authorization：", authorization)
		if authorization == "" {
			log.Println("未包含TOKEN")
			util.Response(c, 500, 500, gin.H{}, "用户未认证")
			c.Abort()
		} else if tokenstring := strings.TrimLeft(authorization, "Bearer"); tokenstring != "" {
			tokenstring = tokenstring[1:]
			_, claims, err := ParseToken(tokenstring)
			if err != nil {
				log.Println("Token 解析出错！")
				log.Println("Token:", tokenstring)
				log.Println(err)
				c.Abort()
			}
			// username := claims.UserName
			// OldIP := claims.IP["ip"]

			cache := common.GetCache()
			redisUtilInstance := util.RedisUtilInstance(cache)

			if isExit, _ := redisUtilInstance.IsBlackList(tokenstring); isExit {
				log.Println("用户%s的TOKEN在黑名单种，拒绝访问", claims.UserName)
				util.Response(c, 500, 500, gin.H{}, "TOKEN in the Blacklist!!!")
				c.Abort()
				return
			}

			//判断TOKEN是否过期
			if isHas, _ := redisUtilInstance.HasKey(tokenstring); isHas {
				expirationTime, _ := redisUtilInstance.HGet(tokenstring, "expirationTime")

				if isEx, _ := TokenIsExpiration(expirationTime); isEx {
					tokenValidTime, _ := redisUtilInstance.GetTokenValidTimeByToken(tokenstring)
					log.Println("Token 作废", tokenstring)
					redisUtilInstance.HSet("blacklist", tokenstring, time.Now().String())

					//compare now and tokenValidTime
					if isOut, _ := TokenIsExpiration(tokenValidTime); isOut {
						//超过有效时间
						log.Println(tokenstring, " already out of valid Time!")
						util.Response(c, 500, 500, gin.H{}, "TOKEN 失效，请重新登录")
						c.Abort()
						return
					} else {
						// Refresh Token: token still in the valid time
						usernameByToken, _ := redisUtilInstance.GetUsernameByToken(tokenstring)
						// ip, _ :=redisUtilInstance.GetIPByToken(tokenstring)
						// username = usernameByToken

						newToken, _ := ReleaseToken(c, model.User{Username: usernameByToken})

						// add new token to redis
						redisUtilInstance.AddTokenTORedis(newToken, usernameByToken, c.Request.RemoteAddr)
						//delete old token
						redisUtilInstance.HDeleteKey(tokenstring)

						tokenstring = newToken
						log.Println("Token still in valid time and has been refreshed!")
						c.Request.Response.Header.Set("Authorization", "Bearer "+newToken)
					}
				}
			}

			// if username != "" && c.Request.Header.Get("Authorization") == authorization {
			// 	if strings.Compare(OldIP, c.Request.RemoteAddr) != 0 {
			// 		log.Println("IP address changed ! valid if IP is in the BlackList !")
			// 		if isExit, _ := redisUtilInstance.IsBlackList(tokenstring); isExit {
			// 			log.Println("用户%s的TOKEN在黑名单种，拒绝访问", claims.UserName)
			// 			util.Response(c, 500, 500, gin.H{}, "TOKEN in the Blacklist!!!")
			// 			c.Abort()
			// 			return
			// 		}

			// 	}

			// }

			c.Set("claims", claims)

		}

	}
}
