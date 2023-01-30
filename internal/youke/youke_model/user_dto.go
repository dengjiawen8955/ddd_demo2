package youke_model

type C2S_SaveYoukeUser struct {
	UserID    string `json:"user_id"`
	IDImg     string `json:"id_img"`
	UserImg   string `json:"user_img"`
	Phone     string `json:"phone"`
	IDNum     string `json:"id_num"`
	Username  string `json:"username"`
	IDAddress string `json:"id_address"`
}

func (y *C2S_SaveYoukeUser) ToDomain() (*YoukeUser, error) {
	UserID, err := NewUserID(y.UserID)
	if err != nil {
		return nil, err
	}
	IDImg, err := NewIDImg(y.IDImg)
	if err != nil {
		return nil, err
	}
	UserImg, err := NewUserImg(y.UserImg)
	if err != nil {
		return nil, err
	}
	Phone, err := NewPhone(y.Phone)
	if err != nil {
		return nil, err
	}
	IDNum, err := NewIDNum(y.IDNum)
	if err != nil {
		return nil, err
	}
	Username, err := NewUsername(y.Username)
	if err != nil {
		return nil, err
	}
	IDAddress, err := NewIDAddress(y.IDAddress)
	if err != nil {
		return nil, err
	}

	return &YoukeUser{
		UserID:    UserID,
		IDImg:     IDImg,
		UserImg:   UserImg,
		Phone:     Phone,
		IDNum:     IDNum,
		Username:  Username,
		IDAddress: IDAddress,
	}, nil
}

type C2S_GetUser struct {
	Phone    string `json:"phone"`
	IDNum    string `json:"id_num"`
	Username string `json:"username"`
}

func (c *C2S_GetUser) ToDomain() (*YoukeUser, error) {
	phone, err := NewPhone(c.Phone)
	if err != nil {
		return nil, err
	}
	idNum, err := NewIDNum(c.IDNum)
	if err != nil {
		return nil, err
	}
	username, err := NewUsername(c.Username)
	if err != nil {
		return nil, err
	}
	return &YoukeUser{
		Phone:    phone,
		IDNum:    idNum,
		Username: username,
	}, nil
}

type S2C_GetUser struct {
	UserID    string `json:"user_id"`
	IDImg     string `json:"id_img"`
	UserImg   string `json:"user_img"`
	Phone     string `json:"phone"`
	IDNum     string `json:"id_num"`
	Username  string `json:"username"`
	IDAddress string `json:"id_address"`
}
