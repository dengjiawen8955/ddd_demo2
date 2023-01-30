package youke_model

import (
	"dc2/internal/common/utils"
)

type OrderPO struct {
	// 主键
	OrderID  int        `column:"order_id" gorm:"primaryKey;autoIncrement"`
	UserID   int        `column:"user_id"`
	CreateAt utils.Time `column:"create_at"`
}

func (OrderPO) TableName() string {
	return "youke_order"
}

func (o *OrderPO) ToDomain() *Order {
	orderID, _ := NewOrderID(utils.IntToString(o.OrderID))
	userID, _ := NewUserID(utils.IntToString(o.UserID))

	return &Order{
		OrderID:  orderID,
		UserID:   userID,
		CreateAt: o.CreateAt,
	}
}

type YoukeUserOrderPO struct {
	OrderId   string     `column:"order_id"`
	UserId    string     `column:"user_id"`
	CreateAt  utils.Time `column:"create_at"`
	IdImg     string     `column:"id_img"`
	UserImg   string     `column:"user_img"`
	Phone     string     `column:"phone"`
	IdNum     string     `column:"id_num"`
	Username  string     `column:"username"`
	IdAddress string     `column:"id_address"`
}

func (r *YoukeUserOrderPO) ToDomain() *Order {
	orderID, _ := NewOrderID(r.OrderId)
	userID, _ := NewUserID(r.UserId)
	createAt := r.CreateAt
	idImg, _ := NewIDImg(r.IdImg)
	userImg, _ := NewUserImg(r.UserImg)
	phone, _ := NewPhone(r.Phone)
	idNum, _ := NewIDNum(r.IdNum)
	username, _ := NewUsername(r.Username)
	idAddress, _ := NewIDAddress(r.IdAddress)
	return &Order{
		OrderID:  orderID,
		UserID:   userID,
		CreateAt: createAt,
		YoukeUser: &YoukeUser{
			UserID:    userID,
			IDImg:     idImg,
			UserImg:   userImg,
			Phone:     phone,
			IDNum:     idNum,
			Username:  username,
			IDAddress: idAddress,
		},
	}
}
