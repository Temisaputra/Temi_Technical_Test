package entity

import (
	"github.com/Temisaputra/warOnk/delivery/presenter"
)

type Voucher struct {
	No          int     `json:"no" gorm:"column:no;primaryKey;autoIncrement"`
	VoucherCode string  `json:"voucher_code" gorm:"column:voucher_code"`
	Discount    float64 `json:"discount" gorm:"column:discount"`
	ExpiredDate string  `json:"expired_date" gorm:"column:expired_date"`
	CreatedAt   string  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   string  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *string `json:"deleted_at" gorm:"column:deleted_at"`
}

func (r *Voucher) TableName() string {
	return "voucher"
}

func (r *Voucher) ToPresenter() presenter.VoucherResponse {
	return presenter.VoucherResponse{
		No:          r.No,
		VoucherCode: r.VoucherCode,
		Discount:    r.Discount,
		ExpiredDate: r.ExpiredDate,
	}
}
