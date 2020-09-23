package web

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-micro-learn/members/resouces"
	"net/http"
)

var (
	ParamsErr = errors.New("param err")
)

func Login(c *gin.Context) {
	username := c.PostForm("user_name")
	pwd := c.PostForm("password")
	if username == "" || pwd == "" {
		c.JSON(http.StatusBadRequest,ParamsErr)
		return
	}
	var uir = resouces.UserInfoResponse{
		Uid:      10086,
		UserName: username,
		Age:      1,
		Sex:      2,
	}
	c.JSON(http.StatusOK, uir)

}
func UserInfo(c *gin.Context) {
	gid := uuid.New()
	uuid := gid.ID()
	var uir = resouces.UserInfoResponse{
		Uid:      uuid,
		UserName: "peanut",
		Age:      1,
		Sex:      2,
	}
	c.JSON(http.StatusOK, uir)
}
