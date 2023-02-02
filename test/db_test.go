package test

import (
	"ddd_demo2/config"
	"ddd_demo2/internal/servers"
	"ddd_demo2/internal/youke/youke_model"
	"fmt"
	"strconv"
	"testing"

	"github.com/dengjiawen8955/stogo/stogo"
	"gorm.io/gorm"
)

var (
	configFile = "../config.yaml"
	db         *gorm.DB
	cfg        *config.SugaredConfig
)

func TestMain(m *testing.M) {
	cfg = config.NewConfig(configFile)

	fmt.Printf("cfg.Mysql.DNS: %v\n", cfg.Mysql.DNS)

	db = servers.NewDB(cfg)

	fmt.Printf("db.Error: %v\n", db.Error)

	m.Run()
}

type YoukeUserPO2 struct {
	UserID    string `column:"user_id"`
	IDImg     string `column:"id_img"`
	UserImg   string `column:"user_img"`
	Phone     string `column:"phone"`
	IDNum     string `column:"id_num"`
	Username  string `column:"username"`
	IDAddress string `column:"id_address"`
}

func (YoukeUserPO2) TableName() string {
	return "youke_user2"
}

func Test_Move_One(t *testing.T) {
	var user youke_model.YoukeUserPO
	err := db.First(&user).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("user: %v\n", user)

	user2 := &YoukeUserPO2{
		UserID:    strconv.Itoa(int(user.UserID)),
		IDImg:     user.IDImg,
		UserImg:   user.UserImg,
		Phone:     user.Phone,
		IDNum:     user.IDNum,
		Username:  user.Username,
		IDAddress: user.IDAddress,
	}

	err = db.Create(user2).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("user: %v\n", user2)
}

func Test_Move_All(t *testing.T) {
	// 读取 youke_user 表然后更新到 youke_user2 表, 主键自增
	var users []*youke_model.YoukeUserPO
	err := db.Find(&users).Error
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("user: %v\n", user)

	var user2s []*YoukeUserPO2
	for i, u1 := range users {
		user2s = append(user2s, &YoukeUserPO2{
			UserID:    fmt.Sprintf("%d", i+1),
			IDImg:     u1.IDImg,
			UserImg:   u1.UserImg,
			Phone:     u1.Phone,
			IDNum:     u1.IDNum,
			Username:  u1.Username,
			IDAddress: u1.IDAddress,
		})
	}

	err = db.Create(&user2s).Error
	if err != nil {
		t.Error(err)
	}
	// fmt.Printf("user: %v\n", user2s)

}

func Test_stogo(t *testing.T) {
	//查询语句
	ssql := `SELECT
	youke_order.*, 
	youke_user2.*
FROM
	youke_user2
	INNER JOIN
	youke_order
	ON 
		youke_user2.user_id = youke_order.user_id`
	//数据Driver
	driver := cfg.Mysql.DNS
	//多连表查询语句直接生成 go DTO 层结构体代码
	stogo.GenerateStruct(ssql, driver)
}

func Test_RawScan(t *testing.T) {
	var result []youke_model.YoukeUserOrderPO
	ssql := `SELECT
	youke_order.*, 
	youke_user2.*
FROM
	youke_user2
	INNER JOIN
	youke_order
	ON 
		youke_user2.user_id = youke_order.user_id`
	err := db.Raw(ssql).Scan(&result).Error
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("result: %v\n", result)
}
