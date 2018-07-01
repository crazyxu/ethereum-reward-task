package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/crazyxu/ethereum-reward-task/model"
)

const ContextUser = "user"
var backed model.BackendApi

func New(backedApi model.BackendApi) *http.Server{
	backed = backedApi
	router := gin.New()
	group := router.Group("/", auth)
	group.POST("/api/eth/tasks", publishTask)
	srv := &http.Server{
		Addr:":8080",
		Handler: router,
	}
	return srv
}

func auth(c *gin.Context){
	key := c.GetHeader("Authorization")
	user, err := queryUser(key)
	if err != nil{
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if user == nil{
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Set(ContextUser, user)
}

//创建一个任务
func publishTask(c *gin.Context){
	type TaskReq struct {
		TemplateId uint `json:"template_id"`
		Name string `json:"name"`
		Description string `json:"description"`
		Reward uint64 `json:"reward"`
		EthPwd string `json:"eth_password"`
	}
	req := TaskReq{}
	if c.Bind(&req) == nil {
		user := getUserInContext(c)
		if user == nil{
			return
		}
		//检查是否有eth账户
		if user.EthAddress == "" {
			c.AbortWithStatus(http.StatusPaymentRequired)
			return
		}
		//检测合约模板
		taskTemp, err := getTaskTemplate(req.TemplateId)
		if err != nil{
			c.AbortWithStatus(http.StatusNotFound)
		}
		//if taskTemp.match(){
		//
		//}

		//发布任务到blockChain
		taskAddr, err := backed.PublishTask(req.Name, req.Description, user.EthAddress, req.Reward)
		if err != nil{
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		//保存任务记录到本地
		id, err := saveTask(req.Name, req.Description, req.Reward, taskAddr)
		if err != nil{
			c.AbortWithError(http.StatusExpectationFailed, err)
			return
		}
		taskResp := Task{
			Id:id,
			Name:req.Name,
			Description:req.Description,
			Reward:req.Reward,
		}
		c.JSON(http.StatusOK, taskResp)
	}
}

func getUserInContext(c *gin.Context) *User{
	v, ok := c.Get(ContextUser)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil
	}
	user, ok := v.(User)
	if !ok{
		c.AbortWithStatus(http.StatusInternalServerError)
		return nil
	}
	return &user
}