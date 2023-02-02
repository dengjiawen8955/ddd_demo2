package user

import (
	"ddd_demo2/internal/user/user_model"
	"errors"

	"gorm.io/gorm"
)

type UserRepo interface {
	Get(*user_model.UserID) (*user_model.User, error)
	Login(*user_model.User) (*user_model.User, error)
	Register(*user_model.User) (*user_model.User, error)
	Save(*user_model.User) (*user_model.User, error)
	// 检查用户是否存在
	Exists(*user_model.UserID) (bool, error)
}

var (
	ErrUserUsernameOrPassword = errors.New("用户名或者密码错误")
	ErrUserNotFound           = errors.New("用户不存在")
)

var _ UserRepo = &MysqlUserRepo{}

type MysqlUserRepo struct {
	db *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) *MysqlUserRepo {
	return &MysqlUserRepo{db: db}
}

func (r *MysqlUserRepo) Login(loginParams *user_model.User) (*user_model.User, error) {
	var userPO user_model.UserPO
	var db = r.db
	var err error

	if loginParams.Username.Value() != "" {
		err = db.Where("username = ? AND password = ?", loginParams.Username.Value(), loginParams.Password.Value()).First(&userPO).Error
	}
	// TODO: 支持其他参数查找

	if err != nil {
		return nil, ErrUserUsernameOrPassword
	}

	return userPO.ToDomain()
}

func (r *MysqlUserRepo) Register(registerParams *user_model.User) (*user_model.User, error) {
	var count int64
	var db = r.db
	var err error

	// 检查是否已经注册
	if registerParams.Username.Value() != "" {
		err = db.Where("username = ?", registerParams.Username.Value()).Count(&count).Error
	}
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, ErrUserAlreadyExists
	}

	// 注册
	user, err := r.Save(registerParams)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MysqlUserRepo) Get(id *user_model.UserID) (*user_model.User, error) {
	var userPO user_model.UserPO
	var db = r.db

	if err := db.Where("id = ?", id.Value()).First(&userPO).Error; err != nil {
		return nil, ErrUserNotFound
	}

	return userPO.ToDomain()
}

func (r *MysqlUserRepo) Save(user *user_model.User) (*user_model.User, error) {
	var userPO = user.ToPO()

	if err := r.db.Save(&userPO).Error; err != nil {
		return nil, err
	}

	return userPO.ToDomain()
}

func (r *MysqlUserRepo) Exists(id *user_model.UserID) (bool, error) {
	var userPO user_model.UserPO
	var db = r.db

	if err := db.Where("id = ?", id.Value()).First(&userPO).Error; err != nil {
		return false, err
	}

	return true, nil
}
