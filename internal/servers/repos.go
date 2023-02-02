package servers

import (
	"ddd_demo2/config"

	"ddd_demo2/internal/bill"
	"ddd_demo2/internal/user"
	"ddd_demo2/internal/youke"

	"github.com/go-redis/redis/v8"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbDriver = "mysql"
)

type Repos struct {
	UserRepo  user.UserRepo
	AuthRepo  user.AuthInterface
	BillRepo  bill.BillRepo
	YoukeRepo youke.YoukeRepo
}

func NewDB(cfg *config.SugaredConfig) *gorm.DB {

	db, err := gorm.Open(mysql.Open(cfg.Mysql.DNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 打印 sql
	db = db.Debug()

	return db
}

func NewCache(cfg *config.SugaredConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       0,
	})
}

func NewRepos(cfg *config.SugaredConfig) *Repos {
	// 持久化类型的 repo
	db := NewDB(cfg)
	userRepo := user.NewMysqlUserRepo(db)
	billRepo := bill.NewMysqlBillRepo(db)
	youkeRepo := youke.NewMysqlYoukeRepo(db)

	// auth 策略
	var authRepo user.AuthInterface
	if cfg.Auth.Active == "redis" {
		authRepo = user.NewRedisAuthRepo(NewCache(cfg), cfg.AuthExpireTime)
	} else {
		authRepo = user.NewJwtAuth(cfg.Auth.PrivateKey, cfg.AuthExpireTime)
	}

	return &Repos{
		UserRepo:  userRepo,
		AuthRepo:  authRepo,
		BillRepo:  billRepo,
		YoukeRepo: youkeRepo,
	}
}
