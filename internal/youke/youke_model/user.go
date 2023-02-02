package youke_model

import (
	"ddd_demo2/internal/common/utils"
)

// user_id, id_img, user_img, phone, id_num, username, id_address

type UserID struct {
	value string
}

func NewUserID(value string) (*UserID, error) {
	return &UserID{
		value: value,
	}, nil
}

func (u *UserID) Value() string {
	if u == nil {
		return ""
	}
	return u.value
}

type IDImg struct {
	value string
}

func NewIDImg(value string) (*IDImg, error) {
	return &IDImg{
		value: value,
	}, nil
}

func (i *IDImg) Value() string {
	if i == nil {
		return ""
	}
	return i.value
}

type UserImg struct {
	value string
}

func NewUserImg(value string) (*UserImg, error) {
	return &UserImg{
		value: value,
	}, nil
}

func (u *UserImg) Value() string {
	if u == nil {
		return ""
	}
	return u.value
}

type Phone struct {
	value string
}

func NewPhone(value string) (*Phone, error) {
	return &Phone{
		value: value,
	}, nil
}

func (p *Phone) Value() string {
	if p == nil {
		return ""
	}
	return p.value
}

type IDNum struct {
	value string
}

func NewIDNum(value string) (*IDNum, error) {
	return &IDNum{
		value: value,
	}, nil
}

func (i *IDNum) Value() string {
	if i == nil {
		return ""
	}
	return i.value
}

type IDAddress struct {
	value string
}

func NewIDAddress(value string) (*IDAddress, error) {
	return &IDAddress{
		value: value,
	}, nil
}

func (i *IDAddress) Value() string {
	if i == nil {
		return ""
	}
	return i.value
}

type Username struct {
	value string
}

func NewUsername(value string) (*Username, error) {
	return &Username{
		value: value,
	}, nil
}

func (u *Username) Value() string {
	if u == nil {
		return ""
	}
	return u.value
}

type YoukeUser struct {
	UserID    *UserID
	IDImg     *IDImg
	UserImg   *UserImg
	Phone     *Phone
	IDNum     *IDNum
	Username  *Username
	IDAddress *IDAddress
}

func (u *YoukeUser) ToPO() *YoukeUserPO {
	userIDInt, _ := utils.StringToInt(u.UserID.Value())
	return &YoukeUserPO{
		UserID:    userIDInt,
		IDImg:     u.IDImg.Value(),
		UserImg:   u.UserImg.Value(),
		Phone:     u.Phone.Value(),
		IDNum:     u.IDNum.Value(),
		Username:  u.Username.Value(),
		IDAddress: u.IDAddress.Value(),
	}
}

func (u *YoukeUser) ToS2C_GetUser() *S2C_GetUser {
	return &S2C_GetUser{
		UserID:    u.UserID.Value(),
		IDImg:     u.IDImg.Value(),
		UserImg:   u.UserImg.Value(),
		Phone:     u.Phone.Value(),
		IDNum:     u.IDNum.Value(),
		Username:  u.Username.Value(),
		IDAddress: u.IDAddress.Value(),
	}
}
