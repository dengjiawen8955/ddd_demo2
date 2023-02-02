package bill

import "ddd_demo2/internal/bill/bill_model"

type BillAppInterface interface {
	CreateBill(bill *bill_model.Bill) error
}

var _ BillAppInterface = &BillApp{}

type BillApp struct {
	BillRepo BillRepo
}

func NewBillApp(billRepo BillRepo) *BillApp {
	return &BillApp{
		BillRepo: billRepo,
	}
}

func (a *BillApp) CreateBill(bill *bill_model.Bill) error {
	return a.BillRepo.Save(bill)
}
