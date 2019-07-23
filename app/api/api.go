package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrum-poker/models"
	"scrum-poker/pkg/e"
	"strconv"
	"math/rand"
	// "fmt"
)

//查
func GetTest(c *gin.Context) {
	str := c.Query("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		code := e.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": "",
		})
		models.CloseDB()
		return
	}
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["model"] = models.GetTestModel(id)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//增
func AddTest(c *gin.Context) {
}

//改
func EditTest(c *gin.Context) {
}

//删
func DeleteTest(c *gin.Context) {
}

//得到链接
func GetUrl(c *gin.Context) {
	randNumber := rand.Int()
	data := make(map[string]interface{})
	data["url"] = "/api/url?code="+strconv.Itoa(randNumber)
	returnjson(c, 1, data)
}
//翻牌
func Poker(c *gin.Context) {
	var spoker models.Scrum_poker
	
	spoker.Url_code = c.Query("url_code")
	spoker.Poker = c.Query("poker")
	user_id, err := c.Cookie("user_id")
	if(err != nil){
		user_id = "1";
	}
	spoker.User_id,_ = strconv.Atoi(user_id)

	where := make(map[string]interface{})
	where["user_id"] = spoker.User_id
	where["Url_code"] = spoker.Url_code
	
	info := models.GetOne(where)
	if(info.Id == 0){
		models.AddPoker(spoker)
	}else{
		models.UpdPoker(where, spoker)
	}
	
	data := make(map[string]interface{})
	data["msg"] = "成功"
	returnjson(c, 1, data)
}
//设置cookier
func SetCk(c *gin.Context) {
	_, err := c.Cookie("user_id")
	if(err != nil){
		randNumber := rand.Intn(10000)
		c.SetCookie("user_id",strconv.Itoa(randNumber) , 3600, "/", "", false, true)
	}
}
//返回值
func returnjson(c *gin.Context, status int, data map[string]interface{}) {
	
	code := e.SUCCESS
	if (status == 0) {
		code = e.ERROR
	}else if (status == 2) {
		code = e.INVALID_PARAMS
	}
	
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}