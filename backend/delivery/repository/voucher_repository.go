package repository

import (
	"context"

	"github.com/Temisaputra/warOnk/delivery/presenter"
	"github.com/Temisaputra/warOnk/delivery/presenter/request"
	"github.com/Temisaputra/warOnk/delivery/presenter/response"
	"github.com/Temisaputra/warOnk/internal/entity"
)

type VoucherRepository interface {
	GetListVoucher(ctx context.Context, params presenter.VoucherParams, pagination *request.Pagination) (res []presenter.VoucherResponse, meta response.Meta, err error)
	GetVoucherByID(ctx context.Context, id int) (presenter.VoucherResponse, error)
	CreateVoucher(ctx context.Context, voucher entity.Voucher) error
	UpdateVoucher(ctx context.Context, voucher entity.Voucher) error
	DeleteVoucher(ctx context.Context, id int) error
	CheckVoucherCode(ctx context.Context, code string) (bool, error)
	ExportVoucher(ctx context.Context) (res []presenter.VoucherResponse, err error)
}
