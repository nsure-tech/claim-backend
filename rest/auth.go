package rest

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"nsure/vote/service"
)

const keyCurrentUser = "__current_user"

func checkToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if len(token) == 0 {
			var err error
			token, err = c.Cookie("accessToken")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusForbidden, newMessageVo(errors.New("token not found")))
				return
			}
		}

		userId, err := service.CheckToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, newMessageVo(err))
			return
		}
		if userId == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, newMessageVo(errors.New("bad token")))
			return
		}

		c.Set(keyCurrentUser, &userId)
		c.Next()
	}
}

func GetCurrentUser(ctx *gin.Context) *string {
	val, found := ctx.Get(keyCurrentUser)
	if !found {
		return nil
	}
	return val.(*string)
}
