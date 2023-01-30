package youke_model

import (
	"dc2/internal/common/utils"
)

type YoukeUserPO struct {
	UserID    int    `column:"user_id" gorm:"primaryKey"`
	IDImg     string `column:"id_img"`
	UserImg   string `column:"user_img"`
	Phone     string `column:"phone"`
	IDNum     string `column:"id_num"`
	Username  string `column:"username"`
	IDAddress string `column:"id_address"`
}

func (YoukeUserPO) TableName() string {
	return "youke_user"
}

func (y *YoukeUserPO) ToDomain() *YoukeUser {
	userIDStr := utils.IntToString(y.UserID)
	UserID, _ := NewUserID(userIDStr)
	IDImg, _ := NewIDImg(y.IDImg)
	UserImg, _ := NewUserImg(y.UserImg)
	Phone, _ := NewPhone(y.Phone)
	IDNum, _ := NewIDNum(y.IDNum)
	Username, _ := NewUsername(y.Username)
	IDAddress, _ := NewIDAddress(y.IDAddress)

	return &YoukeUser{
		UserID:    UserID,
		IDImg:     IDImg,
		UserImg:   UserImg,
		Phone:     Phone,
		IDNum:     IDNum,
		Username:  Username,
		IDAddress: IDAddress,
	}
}
