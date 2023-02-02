package youke

import (
	"ddd_demo2/internal/common/utils"
	"ddd_demo2/internal/youke/youke_model"
	"time"

	"gorm.io/gorm"
)

type YoukeRepo interface {
	// 新增 order 记录
	CreateOrder(userID *youke_model.UserID) error
	GetUserByID(userID *youke_model.UserID) (*youke_model.YoukeUser, error)
	SaveUser(user *youke_model.YoukeUser) (newUser *youke_model.YoukeUser, err error)
	GetUser(getUser *youke_model.YoukeUser) (*youke_model.YoukeUser, error)
	GetOrders(user *youke_model.YoukeUser) ([]*youke_model.Order, error)
}

var _ YoukeRepo = (*MysqlYoukeRepo)(nil)

type MysqlYoukeRepo struct {
	db *gorm.DB
}

func NewMysqlYoukeRepo(db *gorm.DB) YoukeRepo {
	return &MysqlYoukeRepo{db: db}
}

func (r *MysqlYoukeRepo) CreateOrder(userID *youke_model.UserID) error {
	userIDInt, err := utils.StringToInt(userID.Value())
	if err != nil {
		return err
	}

	po := youke_model.OrderPO{
		UserID:   userIDInt,
		CreateAt: utils.Time(time.Now()),
	}

	return r.db.Create(&po).Error
}

func (r *MysqlYoukeRepo) GetUserByID(userID *youke_model.UserID) (*youke_model.YoukeUser, error) {
	var po youke_model.YoukeUserPO
	var db = r.db

	db = db.Where("user_id = ?", userID.Value())

	err := db.First(&po).Error
	if err != nil {
		return nil, ErrUserNotExist
	}

	return po.ToDomain(), nil
}

func (r *MysqlYoukeRepo) SaveUser(user *youke_model.YoukeUser) (newUser *youke_model.YoukeUser, err error) {
	po := user.ToPO()

	err = r.db.Save(po).Error
	if err != nil {
		return nil, err
	}

	return po.ToDomain(), nil
}

func (r *MysqlYoukeRepo) GetUser(getUser *youke_model.YoukeUser) (*youke_model.YoukeUser, error) {
	var db = r.db

	// 通过 phone, id_num, username 不为空的进行模糊查询
	if getUser.Phone != nil {
		db = db.Where("phone like ?", "%"+getUser.Phone.Value()+"%")
	}
	if getUser.IDNum != nil {
		db = db.Where("id_num like ?", "%"+getUser.IDNum.Value()+"%")
	}
	if getUser.Username != nil {
		db = db.Where("username like ?", "%"+getUser.Username.Value()+"%")
	}

	// 根据 id 倒叙
	db = db.Order("user_id desc")

	// 执行查询
	var user youke_model.YoukeUserPO
	err := db.Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}

func (r *MysqlYoukeRepo) GetOrders(user *youke_model.YoukeUser) ([]*youke_model.Order, error) {
	var db = r.db

	// 通过 phone, id_num, username 不为空的进行模糊查询
	if user.Phone != nil {
		db = db.Where("phone like ?", "%"+user.Phone.Value()+"%")
	}
	if user.IDNum != nil {
		db = db.Where("id_num like ?", "%"+user.IDNum.Value()+"%")
	}
	if user.Username != nil {
		db = db.Where("username like ?", "%"+user.Username.Value()+"%")
	}

	// 执行查询
	orderPOs := make([]*youke_model.YoukeUserOrderPO, 0)
	err := db.Table("youke_order").
		Select("youke_order.order_id, youke_order.user_id, youke_order.create_at, youke_user.id_img, youke_user.user_img, youke_user.phone, youke_user.id_num, youke_user.username, youke_user.id_address").
		Joins("left join youke_user on youke_order.user_id = youke_user.user_id").
		Scan(&orderPOs).Error
	if err != nil {
		return nil, err
	}

	// 转换为 order
	returnOrders := make([]*youke_model.Order, 0, len(orderPOs))
	for _, orderPO := range orderPOs {
		returnOrders = append(returnOrders, orderPO.ToDomain())
	}

	return returnOrders, nil
}
