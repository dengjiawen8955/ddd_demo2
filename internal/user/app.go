package user

import (
	"ddd_demo2/internal/bill"
	bill_model "ddd_demo2/internal/bill/bill_model"
	"ddd_demo2/internal/user/user_model"
	"errors"
)

var (
	ErrUserAlreadyExists = errors.New("用户已存在")
)

type UserAppInterface interface {
	Login(loginParams *user_model.User) (*user_model.User, error)
	GetAuthInfo(token string) (*user_model.AuthInfo, error)
	Get(userID *user_model.UserID) (*user_model.User, error)
	Register(registerParams *user_model.User) (*user_model.User, error)
	Transfer(fromUserID, toUserID *user_model.UserID, amount *user_model.Amount, toCurrency *user_model.Currency) error
}

var _ UserAppInterface = &UserApp{}

type UserApp struct {
	userRepo        UserRepo
	authRepo        AuthInterface
	transferService TransferService
	rateService     RateService
	billApp         bill.BillAppInterface
}

func NewUserApp(userRepo UserRepo, authRepo AuthInterface, billRepo bill.BillRepo) UserAppInterface {
	return &UserApp{
		userRepo:        userRepo,
		authRepo:        authRepo,
		transferService: NewTransferService(),
		rateService:     NewRateService(),
		billApp:         bill.NewBillApp(billRepo),
	}
}

func (u *UserApp) Login(loginParams *user_model.User) (*user_model.User, error) {
	// 登录
	user, err := u.userRepo.Login(loginParams)
	if err != nil {
		return nil, err
	}

	// 生成 token
	authInfo := &user_model.AuthInfo{
		UserID: user.ID.Value(),
	}
	token, err := u.authRepo.Set(authInfo)
	if err != nil {
		return nil, err
	}

	user.Token = token

	return user, nil
}

// GetAuthInfo 从 token 中获取用户信息
func (u *UserApp) GetAuthInfo(token string) (*user_model.AuthInfo, error) {
	return u.authRepo.Get(token)
}

// Get 获取用户信息
func (u *UserApp) Get(userID *user_model.UserID) (*user_model.User, error) {
	user, err := u.userRepo.Get(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserApp) Register(registerParams *user_model.User) (*user_model.User, error) {
	user, err := u.userRepo.Register(registerParams)
	if user != nil || err == nil {
		return nil, ErrUserAlreadyExists
	}

	// 生成 token
	authInfo := &user_model.AuthInfo{
		UserID: user.ID.Value(),
	}
	token, err := u.authRepo.Set(authInfo)
	if err != nil {
		return nil, err
	}

	user.Token = token

	return user, nil
}

func (u *UserApp) Transfer(fromUserID, toUserID *user_model.UserID, amount *user_model.Amount, toCurrency *user_model.Currency) error {
	// 读数据
	fromUser, err := u.userRepo.Get(fromUserID)
	if err != nil {
		return err
	}

	toUser, err := u.userRepo.Get(toUserID)
	if err != nil {
		return err
	}

	rate, err := u.rateService.GetRate(fromUser.Currency, toCurrency)
	if err != nil {
		return err
	}

	// 转账
	err = u.transferService.Transfer(fromUser, toUser, amount, rate)
	if err != nil {
		return err
	}

	// 保存数据
	u.userRepo.Save(fromUser)
	u.userRepo.Save(toUser)

	// 保存账单
	bill := &bill_model.Bill{
		FromUserID: fromUser.ID,
		ToUserID:   toUser.ID,
		Amount:     amount,
		Currency:   toCurrency,
	}
	err = u.billApp.CreateBill(bill)
	if err != nil {
		return err
	}

	return nil
}
