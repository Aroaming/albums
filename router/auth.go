package router

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/orange-jacky/albums/data"
	"github.com/orange-jacky/albums/util"
	"time"
)

func GetAuthMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			user := util.GetUser()
			err := user.CheckUser(username)
			if err != nil {
				return username, false
			}
			if err != nil {
				return username, false
			}
			return username, true
		},
		Authorizator: func(username string, c *gin.Context) bool {
			mylog := util.GetMylog()
			user := util.GetUser()
			if err := user.CheckUser(username); err != nil {
				mylog.Infof("Authorizator [%v] fail,[%v]", username, err)
				return false
			}
			mylog.Infof("Authorizator [%v] success", username)
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			resp := data.Response{Status: code, Data: message}
			c.JSON(code, resp)
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",
		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}
}
