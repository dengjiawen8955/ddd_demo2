package user

import (
	"ddd_demo2/internal/common/logs"
	"ddd_demo2/internal/servers/web/response"
	"ddd_demo2/internal/user/user_model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type UserHandler struct {
	UserApp UserAppInterface
}

func NewUserHandler(userApp UserAppInterface) *UserHandler {
	return &UserHandler{
		UserApp: userApp,
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	var err error
	req := &user_model.C2S_Login{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 转化为领域对象 + 参数验证
	loginParams, err := req.ToDomain()

	// 调用应用层
	user, err := u.UserApp.Login(loginParams)
	if err != nil {
		logs.Errorf("[Login] failed, err: %w", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	// domain 转为 s2c
	s2c := user.ToS2C_Login()

	response.Ok(c, s2c)
}

// UserInfo 获取用户信息
func (u *UserHandler) UserInfo(c *gin.Context) {
	userIDStr := c.GetString(UserIDKey)

	userID, err := user_model.NewUserID(userIDStr)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := u.UserApp.Get(userID)
	if err != nil {
		logs.Errorf("[UserInfo] failed, err: %w", err)
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	s2c := user.ToS2C_UserInfo()

	response.Ok(c, s2c)
}

// Register 注册
func (u *UserHandler) Register(c *gin.Context) {
	var err error
	req := &user_model.C2S_Register{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 转化为领域对象 + 参数验证
	registerParams, err := req.ToDomain()
	if err != nil {
		logs.Errorf("[Register] failed, err: %w", err)
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用应用层
	user, err := u.UserApp.Register(registerParams)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	s2c := user.ToS2C_Login()

	response.Ok(c, s2c)
}

// Transfer 转账
func (u *UserHandler) Transfer(c *gin.Context) {
	var err error
	req := &user_model.C2S_Transfer{}

	// 解析参数
	if err = c.ShouldBindJSON(req); err != nil {
		response.Err(c, http.StatusBadRequest, err.Error())
		return
	}

	// 转化为领域对象 + 参数验证
	fromUserID, err := user_model.NewUserID(c.GetString(UserIDKey))
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}
	toUserID, err := user_model.NewUserID(req.ToUserID)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	amountDecimal, err := decimal.NewFromString(req.Amount)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	amount, err := user_model.NewAmount(amountDecimal)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	toCurrency, err := user_model.NewCurrency(req.Currency)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 调用应用层
	err = u.UserApp.Transfer(fromUserID, toUserID, amount, toCurrency)
	if err != nil {
		response.Err(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Ok(c)
}
