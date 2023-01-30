package youke

import (
	"dc2/internal/youke/youke_model"
	"errors"
)

var (
	ErrUserNotExist = errors.New("user not exist")
)

type YoukeAppInterface interface {
	// 新增 order 记录
	CreateOrder(userID *youke_model.UserID) error
	// 连表查询 order 记录
	GetOrders(user *youke_model.YoukeUser) ([]*youke_model.Order, error)
	// 查询用户信息
	GetUser(getUser *youke_model.YoukeUser) (*youke_model.YoukeUser, error)
	// 新增或者更新用户信息
	SaveUser(saveUser *youke_model.YoukeUser) (*youke_model.YoukeUser, error)
}

var _ YoukeAppInterface = (*YoukeApp)(nil)

type YoukeApp struct {
	youkeRepo YoukeRepo
}

func NewYoukeApp(youkeRepo YoukeRepo) *YoukeApp {
	return &YoukeApp{
		youkeRepo: youkeRepo,
	}
}

func (a *YoukeApp) CreateOrder(userID *youke_model.UserID) error {
	// 检查用户是否存在, 保证订单和用户对应
	user, err := a.youkeRepo.GetUserByID(userID)
	if err != nil {
		return err
	}

	return a.youkeRepo.CreateOrder(user.UserID)
}

func (a *YoukeApp) GetUser(getUser *youke_model.YoukeUser) (*youke_model.YoukeUser, error) {
	user, err := a.youkeRepo.GetUser(getUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *YoukeApp) SaveUser(saveUser *youke_model.YoukeUser) (*youke_model.YoukeUser, error) {
	return a.youkeRepo.SaveUser(saveUser)
}

func (a *YoukeApp) GetOrders(user *youke_model.YoukeUser) ([]*youke_model.Order, error) {
	orders, err := a.youkeRepo.GetOrders(user)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
