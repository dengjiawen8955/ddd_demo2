package bill

import (
	"ddd_demo2/internal/bill/bill_model"

	"gorm.io/gorm"
)

type BillRepo interface {
	Save(bill *bill_model.Bill) error
}

type MysqlBillRepo struct {
	db *gorm.DB
}

func NewMysqlBillRepo(db *gorm.DB) *MysqlBillRepo {
	return &MysqlBillRepo{db: db}
}

func (r *MysqlBillRepo) Save(bill *bill_model.Bill) error {
	billPO := bill.ToPO()
	return r.db.Save(billPO).Error
}
