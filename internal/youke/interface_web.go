package youke

import (
	"ddd_demo2/internal/common/logs"
	"ddd_demo2/internal/servers/web/response"
	"ddd_demo2/internal/youke/youke_model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type YoukeHandler struct {
	youkeApp YoukeAppInterface
}

func NewYoukeHandler(youkeApp YoukeAppInterface) *YoukeHandler {
	return &YoukeHandler{youkeApp: youkeApp}
}

// 创建订单
func (h *YoukeHandler) CreateOrder(c *gin.Context) {
	var err error
	req := &youke_model.C2S_CreateOrder{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}
	logs.Debugf("[YoukeHandler] [CreateOrder] req=%#v", req)

	// 转化为领域对象 + 参数验证
	userID, err := youke_model.NewUserID(req.UserID)
	if err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	err = h.youkeApp.CreateOrder(userID)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, nil)
}

// SaveUser 创建或者更新用户
func (h *YoukeHandler) SaveUser(c *gin.Context) {
	var err error
	req := &youke_model.C2S_SaveYoukeUser{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}
	logs.Debugf("[YoukeHandler] [SaveUser] req=%#v", req)

	// 转化为领域对象 + 参数验证
	user, err := req.ToDomain()
	if err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	_, err = h.youkeApp.SaveUser(user)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c, nil)
}

// GetUser 花式查询用户
func (h *YoukeHandler) GetUser(c *gin.Context) {
	var err error
	req := &youke_model.C2S_GetUser{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}
	logs.Debugf("[YoukeHandler] [GetUser] req=%#v", req)

	// 转化为领域对象 + 参数验证
	user, err := req.ToDomain()
	if err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	user, err = h.youkeApp.GetUser(user)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	s2c := user.ToS2C_GetUser()

	response.Ok(c, s2c)
}

// GetUserOrders 查询订单&用户
func (h *YoukeHandler) GetUserOrders(c *gin.Context) {
	var err error
	req := &youke_model.C2S_GetUserOrders{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}
	logs.Debugf("[YoukeHandler] [GetUserOrders] req=%#v", req)

	// 转化为领域对象 + 参数验证
	userParams, err := req.ToDomain()
	if err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	orders, err := h.youkeApp.GetOrders(userParams)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	s2c := make([]*youke_model.S2C_GetUserOrder, 0, len(orders))
	for _, order := range orders {
		s2c = append(s2c, order.ToS2C_GetUserOrder())
	}

	response.Ok(c, s2c)
}
