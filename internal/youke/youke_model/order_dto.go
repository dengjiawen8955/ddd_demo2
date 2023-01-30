package youke_model

import "dc2/internal/common/utils"

type C2S_CreateOrder struct {
	UserID string `json:"user_id"`
}

type S2C_GetUserOrder struct {
	OrderId   string     `json:"order_id"`
	UserId    string     `json:"user_id"`
	CreateAt  utils.Time `json:"create_at"`
	IdImg     string     `json:"id_img"`
	UserImg   string     `json:"user_img"`
	Phone     string     `json:"phone"`
	IdNum     string     `json:"id_num"`
	Username  string     `json:"username"`
	IdAddress string     `json:"id_address"`
}

type C2S_GetUserOrders struct {
	Phone    string `json:"phone"`
	IdNum    string `json:"id_num"`
	Username string `json:"username"`
}

func (y *C2S_GetUserOrders) ToDomain() (*YoukeUser, error) {
	Phone, err := NewPhone(y.Phone)
	if err != nil {
		return nil, err
	}
	IDNum, err := NewIDNum(y.IdNum)
	if err != nil {
		return nil, err
	}
	Username, err := NewUsername(y.Username)
	if err != nil {
		return nil, err
	}

	return &YoukeUser{
		Phone:    Phone,
		IDNum:    IDNum,
		Username: Username,
	}, nil
}
