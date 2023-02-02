package youke_model

import (
	"ddd_demo2/internal/common/utils"
)

type OrderID struct {
	value string
}

func NewOrderID(value string) (*OrderID, error) {
	return &OrderID{
		value: value,
	}, nil
}

func (o *OrderID) Value() string {
	if o == nil {
		return ""
	}
	return o.value
}

type Order struct {
	OrderID   *OrderID
	UserID    *UserID
	CreateAt  utils.Time
	YoukeUser *YoukeUser
}

func (o *Order) ToPO() *OrderPO {
	orderIDInt, _ := utils.StringToInt(o.OrderID.Value())
	userIDInt, _ := utils.StringToInt(o.UserID.Value())
	return &OrderPO{
		OrderID:  orderIDInt,
		UserID:   userIDInt,
		CreateAt: o.CreateAt,
	}
}

func (o *Order) ToS2C_GetUserOrder() *S2C_GetUserOrder {
	return &S2C_GetUserOrder{
		OrderId:   o.OrderID.Value(),
		UserId:    o.UserID.Value(),
		CreateAt:  o.CreateAt,
		IdImg:     o.YoukeUser.IDImg.Value(),
		UserImg:   o.YoukeUser.UserImg.Value(),
		Phone:     o.YoukeUser.Phone.Value(),
		IdNum:     o.YoukeUser.IDNum.Value(),
		Username:  o.YoukeUser.Username.Value(),
		IdAddress: o.YoukeUser.IDAddress.Value(),
	}
}
