package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrum-poker/models"
	"scrum-poker/pkg/e"
	"strconv"
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
