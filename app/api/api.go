package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrum-poker/models"
	"scrum-poker/pkg/e"
	"strconv"
	"strings"
	"math/rand"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
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
//websocket
var upGrader = websocket.Upgrader{  
	CheckOrigin: func (r *http.Request) bool {  
	   return true  
	},  
} 
type WsPoker struct {
	ws *websocket.Conn
	Poker string `json:"poker"`
	code string
	mt int
}
var wspoker []WsPoker
//webSocket请求ping 返回pong  
func Ping(c *gin.Context) {  
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)  
	if err != nil {  
	   return  
	}  
	defer ws.Close()  
	code := ""
	for {
	   //读取ws中的数据  
	   mt, message, err := ws.ReadMessage()  
	   if err != nil {  
		  break  
	   }
	   msg := string(message)
	     
	   //链接
	   if (strings.Index(msg, "type=1") != -1) {
			fmt.Println(wspoker)
			str := strings.Split(msg, "&")
			strr := strings.Split(str[1], "=")
			code = strr[1]
			wsClint := WsPoker{ws:ws,code:strr[1],mt:mt}
			wspoker = append(wspoker, wsClint)
			fmt.Println(wspoker)
		//翻牌
	   } else if (strings.Index(msg, "type=2") != -1) {
		    len := len(msg)
			poker := msg[len-1:]
			isOver := 1
			num := 0
			for  k,v := range wspoker{
				if (v.ws == ws){
					wspoker[k].Poker = poker
				}
				if( v.code == code){
					if(wspoker[k].Poker == ""){
						isOver = 0
					}
					num += 1
				}
				
			}
			fmt.Println(wspoker)
			fmt.Println(isOver)
			fmt.Println(num)
			if(isOver == 1 && num != 1){
				
				res := make(map[string]int)
				for  _,v := range wspoker{
					if (v.code == code){
						if _, ok := res[v.Poker]; ok {
							res[v.Poker] += 1
						}else{
							res[v.Poker] = 1
						}
					}
				}
				mjson,_ :=json.Marshal(res)
				//写入ws数据
				for  _,v := range wspoker{
					if (v.code == code){
						err = v.ws.WriteMessage(v.mt, mjson)
						fmt.Println(v.mt)
						fmt.Println("发送消息")  
						if err != nil {  
							
						} 
					}
				} 
			}
		//重新开始	
	   } else if (strings.Index(msg, "type=3") != -1) {
			for  k,v := range wspoker{
				if(v.code == code){
					wspoker[k].Poker = ""
				}
			}
			fmt.Println(wspoker)
	   }
	   
	}
	//断开链接
	for  k,v := range wspoker{
		if (v.ws == ws){
			wspoker = append(wspoker[:k], wspoker[k+1:]...)
		}
	}
	fmt.Println(wspoker)
 }  
