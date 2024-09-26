package middlewares

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NormalUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// User can only edit their own blog
		OwnerID, err := strconv.Atoi(c.Query("userID"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Authentication Failed 1",
			})
			c.Abort()
			return
		}

		reqUserID := c.GetInt("request_userid")
		if reqUserID == 0 || reqUserID != OwnerID {
			c.JSON(http.StatusBadRequest, gin.H{
				"result": "Authentication Failed 2",
			})
			c.Abort()
			return
		}

		// tmp := &struct {
		// 	Uid int `json:"Uid"`
		// }{}
		// err = c.ShouldBindJSON(tmp)
		// if err == nil && tmp.Uid != reqUserID {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"result": "Authentication Failed 3",
		// 	})
		// 	c.Abort()
		// 	return
		// }

	}
}
